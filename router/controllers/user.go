package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService  // 用户服务
	minioServer *services.MinioService // MinIO 服务
}

// NewUserController 创建一个新的 UserController 实例
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register 处理用户注册请求
func (ctrl *UserController) Register(c *gin.Context) {
	var ur services.UserRegisterService

	//解析数据
	if err := c.ShouldBind(&ur); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	//todo 校验数据
	//if err := ctrl.userService.Struct(ur); err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(120, "数据验证失败"))
	//	return
	//}

	//数据库记录用户信息
	user, err := ctrl.userService.Register(c, ur)
	if err != nil {
		ErrorResponse(c, 1, "注册用户失败", err)
	}

	//初始化名字为为用户ID的MINIO 存储桶
	if err := ctrl.minioServer.InitUserMinio(c, user); err != nil {
		ErrorResponse(c, 1, "初始化用户存储失败", err)
	}
	SuccessResponse(c, 1, "用户注册成功", nil)
}

// Login 处理用户登录请求
func (ctrl *UserController) Login(c *gin.Context) {
	var ul services.UserLoginService
	if err := c.ShouldBind(&ul); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	//todo 数据校验
	//if err := ctrl.V.Struct(userLogin); err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "登录信息验证失败"))
	//	return
	//}

	attempts, err := ctrl.userService.CheckLoginAttempts(c, ul.Email)
	if err != nil {
		ErrorResponse(c, 1, "登录失败", err)
	}

	if !attempts {
		ErrorResponse(c, 1, "登录重试次数过多,请稍后重试", err)
	}

	token, err := ctrl.userService.Login(c, ul)
	if err != nil {
		ctrl.userService.IncrementLoginAttempts(c, ul.Email)
		ErrorResponse(c, 1, "登录失败", err)
	}

	ctrl.userService.ResetLoginAttempts(c, ul.Email)
	SuccessResponse(c, 1, "用户登录成功", token)
}

// UpdatePassword 处理用户更新密码请求
func (ctrl *UserController) UpdatePassword(c *gin.Context) {
	var up services.UpdatePasswordService

	if err := c.ShouldBindJSON(&up); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	if err := ctrl.userService.UpdatePassword(c, up); err != nil {
		ErrorResponse(c, 1, "密码更新失败", err)
	}

	//todo jwt失效
	SuccessResponse(c, 1, "密码更新成功", nil)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {

}

//func (u *User) UpdateProfile(c *gin.Context) {
//
//}
//
//func (u *User) JWTAuthenticate(c *gin.Context) {
//
//	uid, err := u.JWT.Parse(c, c.GetHeader("jwt"))
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "JWT 解析失败", nil))
//		return
//	}
//
//	getUser, err := u.User.GetUser(c, uid)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	if getUser.ID == "" {
//		c.JSON(http.StatusBadRequest, "或者用户不存在")
//		return
//	}
//
//	gen, err := u.JWT.Gen(c, getUser.ID)
//	if err != nil {
//		c.JSON(http.StatusBadGateway, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, NewApiJsonResult(200, "登录成功", res.Login{
//		UserID: getUser.ID,
//		Token:  gen,
//	}))
//	log.Println(getUser.ID, gen)
//}
