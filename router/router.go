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
func InitRoute(conf *types.Config, user *controllers.UserController, auth *middlewares.Auth,
	file *controllers.FileController, folder *controllers.FolderController, ws *controllers.WebsocketController) *gin.Engine {

	// 根据开发模式初始化路由
	router := NewRouter(conf.Develop.DevMod)

	// 定义未找到路由的处理
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	// 创建 API 路由组
	api := router.Group("/api")
	// 创建需要认证的 API 路由组，并使用认证中间件
	authAPI := api.Group("").Use(auth.VerifyUserToken())

	{ // 用户相关路由
		api.POST("/login", user.Login)                       // 登录
		api.POST("/register", user.Register)                 // 注册
		authAPI.PATCH("/user_password", user.UpdatePassword) // 更新密码
	}

	{ // 文件和文件夹相关路由
		authAPI.GET("/file/:file-id", file.GetFileMeta)            // 获取文件元数据
		authAPI.GET("/file/:file-id/content", file.GetFileContent) // 获取文件内容
		authAPI.POST("/file", file.CreateFile)                     // 创建文件
		authAPI.PUT("/file", file.UpdateFile)                      // 更新文件
		authAPI.DELETE("/file/:file-id", file.DeleteFile)          // 删除文件
		authAPI.PATCH("/file", file.RenameFile)                    // 重命名文件

		authAPI.GET("/folder", folder.ListFolder)                 // 获取文件夹列表
		authAPI.POST("/folder", folder.CreatFolder)               // 创建文件夹
		authAPI.DELETE("/folder/:folder-id", folder.DeleteFolder) // 删除文件夹
		authAPI.PATCH("/folder", folder.RenameFolder)             // 重命名文件夹
	}

	{ // WebSocket 相关路由
		api.GET("/websocket/connect", ws.WebsocketConn) // 建立 WebSocket 连接
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
