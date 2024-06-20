package router

import (
	"fsm/middlewares"
	"fsm/pkg/types"
	"fsm/router/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRoute 初始化路由
// 参数：配置文件、用户控制器、认证中间件、文件控制器、文件夹控制器、WebSocket 控制器
func InitRoute(conf *types.Config, userController *controllers.UserController, authMiddlewares *middlewares.Auth,
	fileController *controllers.FileController, folderController *controllers.FolderController, wsController *controllers.WebsocketController) *gin.Engine {

	// 根据开发模式初始化路由
	router := NewRouter(conf.Develop.DevMod)

	// 定义未找到路由的处理
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	// 创建 API 路由组
	api := router.Group("/api")

	//v1 版本路由
	v1 := router.Group("/v1")
	v1.Use(authMiddlewares.VerifyUserToken())

	// 用户相关路由
	user := v1.Group("user")
	{
		api.POST("/login", userController.Login)       // 登录
		api.POST("/register", userController.Register) // 注册
		user.GET("")
		user.DELETE("", userController.DeleteUser)
		user.PATCH("", userController.UpdatePassword) // 更新密码
	}

	// 文件相关路由
	file := v1.Group("file")
	{
		file.GET(":file-id/content", fileController.GetFileContent) // 获取文件内容
		file.GET(":file-id", fileController.GetFileMeta)            // 获取文件元数据
		file.POST("", fileController.CreateFile)                    // 创建文件
		file.PUT(":file-id", fileController.UpdateFile)             // 更新文件元信息
		//file.PUT(":file-id/content", fileController.GetFileContent) // 获取文件内容
		file.DELETE(":file-id", fileController.DeleteFile) // 删除文件
	}

	//文件夹相关路由
	folder := v1.Group("folder")
	{
		folder.GET("", folderController.ListFolder) // 获取文件夹列表
		//folder.GET(":folder-id",folderController.)          //获取当前文件夹的信息（子文件数量和文件大小）
		folder.POST("", folderController.CreatFolder)              // 创建文件夹
		folder.DELETE(":folder-id", folderController.DeleteFolder) // 删除文件夹
		folder.PUT(":folder-id", folderController.RenameFolder)    // 更新文件夹
	}

	// WebSocket 相关路由
	ws := v1.Group("ws")
	{
		ws.GET("", wsController.WebsocketConn) // 建立 WebSocket 连接
	}

	return router
}

// NewRouter 创建一个新的 Gin 引擎实例
// 参数：开发模式标志
func NewRouter(devmod bool) *gin.Engine {
	if devmod {
		return gin.Default() // 开发模式下使用默认的 Gin 引擎
	}
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
