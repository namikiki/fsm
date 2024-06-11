package handle

import (
	"fsm/pkg/jwt"
	"log"
	"net/http"

	"fsm/api/req"
	"fsm/api/res"
	"fsm/pkg/sync"
	"fsm/pkg/types"
	"fsm/pkg/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"github.com/minio/minio-go/v7"
)

type User struct {
	V    *validator.Validate
	mio  *minio.Client
	Sync *sync.Syncer
	User *user.Service
	JWT  jwt.Service
}

func NewUser(v *validator.Validate, mio *minio.Client, sync *sync.Syncer, user *user.Service, jwt jwt.Service) User {
	return User{
		V:    v,
		mio:  mio,
		Sync: sync,
		User: user,
		JWT:  jwt,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (u *User) WebsocketConn(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(1003, NewErrorApiResult(501, err.Error()))
		return
	}

	wsClient := types.SyncClient{
		UserID:   c.GetHeader("userID"),
		ClientID: c.GetHeader("clientID"),
		Conn:     conn,
	}

	u.Sync.WebsocketConnChannel <- wsClient
}

// Register 用户注册
func (u *User) Register(c *gin.Context) {
	var ur req.UserRegister
	if err := c.ShouldBind(&ur); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	if err := u.V.Struct(ur); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "数据验证失败"))
		return
	}

	uu, err := u.User.Register(c, ur)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "注册用户失败"))
		return
	}

	if err := u.mio.MakeBucket(c, uu.ID, minio.MakeBucketOptions{ObjectLocking: false}); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建用户存储失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(201, "注册成功", nil))

	//if err := u.mio.EnableVersioning(c, user.ID); err != nil {
	//	log.Printf("init user minio :%v", err)
	//	return
	//}

}

// Login 用户使用密码登录
func (u *User) Login(c *gin.Context) {
	//session := sessions.Default(c)
	var userLogin req.UserLogin

	if err := c.ShouldBind(&userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析登录信息失败"))
		return
	}

	if err := u.V.Struct(userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "登录信息验证失败"))
		return
	}

	userID, token, err := u.User.Login(c, userLogin.Email, userLogin.PassWord)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "登录失败"))
		return
	}

	log.Println("用户", userID, "登录成功")
	c.JSON(http.StatusOK, NewApiJsonResult(200, "登录成功", res.Login{
		Token:  token,
		UserID: userID,
	}))
	//	c.JSON(http.StatusOK, NewApiResult(200, "登录成功", res.Login{
	//	Token:  token,
	//	UserID: userID,
	//}))

	//session.Set("userid", userLogin.ID)
	//session.Save()
}

// UpdatePassword 用户更新密码
func (u *User) UpdatePassword(c *gin.Context) {
	var up req.UpdatePassword
	if err := c.ShouldBindJSON(&up); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析修改密码信息失败"))
		return
	}

	if err := u.User.UpdatePassword(c, up); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "更新密码失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "密码更新成功", nil))
}

func (u *User) Delete(c *gin.Context) {
	session := sessions.Default(c)

	uid := session.Get("userid").(string)
	session.Clear()
	c.JSON(200, gin.H{
		"code": 100,
		"data": res.UserLogin{
			Uid: uid,
		}})
}

func (u *User) UpdateProfile(c *gin.Context) {

}

func (u *User) JWTAuthenticate(c *gin.Context) {

	uid, err := u.JWT.Parse(c, c.GetHeader("jwt"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "JWT 解析失败", nil))
		return
	}

	getUser, err := u.User.GetUser(c, uid)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if getUser.ID == "" {
		c.JSON(http.StatusBadRequest, "或者用户不存在")
		return
	}

	gen, err := u.JWT.Gen(c, getUser.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(200, "登录成功", res.Login{
		UserID: getUser.ID,
		Token:  gen,
	}))
	log.Println(getUser.ID, gen)
}