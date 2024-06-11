package router

import (
	"fsm/middlewares"
	"fsm/pkg/types"
	"fsm/router/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(conf *types.Config, user *controllers.UserController, auth *middlewares.Auth,
	fs *controllers.FSController, ws *controllers.WebsocketController) *gin.Engine {

	router := NewRouter(conf.Develop.DevMod)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	api := router.Group("/api")
	authAPI := api.Group("").Use(auth.VerifyUserToken())

	{ //user
		api.POST("/login", user.Login)
		api.POST("/register", user.Register)
		authAPI.PATCH("/user_password", user.UpdatePassword)
	}

	{ //file folder
		authAPI.GET("/file/:file-id", fs.GetFileMeta)
		authAPI.GET("/file/:file-id/content", fs.GetFileContent)
		authAPI.POST("/file", fs.CreateFile)
		authAPI.PUT("/file", fs.UpdateFile)
		authAPI.DELETE("/file/:file-id", fs.DeleteFile)
		authAPI.PATCH("/file", fs.RenameFile)

		authAPI.GET("/folder", fs.ListFolder)
		authAPI.POST("/folder", fs.CreatFolder)
		authAPI.DELETE("/folder/:folder-id", fs.DeleteFolder)
		authAPI.PATCH("/folder", fs.RenameFolder)
	}

	{ //websocket
		api.GET("/websocket/connect", ws.WebsocketConn)
	}

	return router
}

func NewRouter(devmod bool) *gin.Engine {
	if devmod {
		return gin.Default()
	}
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
