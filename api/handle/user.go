package handle

import (
	"log"

	"fsm/api/req"
	"fsm/api/res"
	"fsm/pkg/domain"
	"fsm/pkg/sync"
	"fsm/pkg/types"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"github.com/minio/minio-go/v7"
)

type User struct {
	domain.UserService
	V    *validator.Validate
	mio  *minio.Client
	sync *sync.Syncer
}

func NewUser(dus domain.UserService, v *validator.Validate, mio *minio.Client, sync *sync.Syncer) User {
	return User{
		UserService: dus,
		V:           v,
		mio:         mio,
		sync:        sync,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (u *User) WebsocketConn(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	client := types.SyncClient{
		UserID:   c.Query("uid"),
		ClientID: c.Query("cid"),
		Conn:     conn,
	}

	u.sync.WebsocketConnChannel <- client
}

func (u *User) Register(c *gin.Context) {
	session := sessions.Default(c)
	var ur req.UserRegister
	if err := c.ShouldBind(&ur); err != nil {
		log.Println("bind err") //数据类型错误
		return
	}

	if err := u.V.Struct(ur); err != nil {
		log.Println("vali err") //数据不能为空
		return
	}

	user, err := u.UserService.Register(c, ur.Email, ur.PassWord, ur.UserName)
	if err != nil {
		log.Println("store err") //数据不能为空
		return
	}

	if err := u.mio.MakeBucket(c, user.ID, minio.MakeBucketOptions{ObjectLocking: false}); err != nil {
		log.Printf("init user minio :%v", err)
		return
	}

	if err := u.mio.EnableVersioning(c, user.ID); err != nil {
		log.Printf("init user minio :%v", err)
		return
	}

	session.Set("userid", user.ID)
	session.Save()

	c.JSON(200, gin.H{
		"code": 100,
		"data": res.UserLogin{
			Uid: user.ID,
		}})

}

func (u *User) Login(c *gin.Context) {
	session := sessions.Default(c)
	var ulogin req.UserLogin

	if err := c.ShouldBind(&ulogin); err != nil {
		log.Println("bind err")
		return
	}

	log.Println(ulogin)

	if err := u.V.Struct(ulogin); err != nil {
		log.Println("vali err")
		return
	}

	user, err := u.UserService.Login(c, ulogin.Email, ulogin.PassWord)
	if err != nil {
		log.Printf("Login err %v", err)
		return
	}

	session.Set("userid", user.ID)
	session.Save()

	c.JSON(200, gin.H{
		"uid": user.ID,
	})
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

func (u *User) Update(c *gin.Context) {

}

func (u *User) Logout(c *gin.Context) {

}
