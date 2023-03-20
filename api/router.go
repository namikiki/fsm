package api

import (
	"log"
	"net/http"

	"fsm/api/handle"
	"fsm/pkg/types"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, user handle.User, file handle.File, dir handle.Dir, syncTask handle.SyncTask, common handle.Common) {
	//store := cookie.NewStore([]byte("secret"))
	//app.Use(sessions.Sessions("sessionId", store))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	app := router.Group("")
	app.Use(common.VerifyUserToken())

	{ // test
		app.GET("/test", func(c *gin.Context) {
			log.Println(c.GetHeader("userID"))
			log.Println(c.GetHeader("clientID"))
			c.JSON(http.StatusNotFound, "123123")
		})
	}

	{ //user
		app.GET("/websocket/connect", user.WebsocketConn)

		router.POST("/login", user.Login)
		router.POST("/register", user.Register)
		app.DELETE("/user", user.Delete)
	}

	{ // file
		app.GET("/file/:fileID", file.Open)
		app.POST("/file", file.Create)
		app.DELETE("/file", file.Delete)
		app.PUT("/file", file.Update)

		app.PUT("/file/name", file.Rename)
		app.GET("/files/:syncID", file.GetAllFileBySyncID)
		app.GET("/file/metadata", file.GetMetadata)
	}

	{ // dir
		app.POST("/dir", dir.Create)
		app.DELETE("/dir", dir.Delete)
		app.GET("/dir", dir.ReadDir)

		app.PUT("/dir/name", dir.Rename)
		app.GET("/dirs/path", dir.GetAllDirByPath)
		app.GET("/dirs/sid/:syncID", dir.GetAllDirBySyncID)
	}

	{ //syncTask
		app.GET("/synctask/:syncID", syncTask.Get)
		app.POST("/synctask", syncTask.Create)
		app.DELETE("/synctask/:syncID", syncTask.Delete)

		app.GET("/synctasks", syncTask.GetAllSyncTask)
	}
}

func New(conf *types.Config) *gin.Engine {
	if conf.Develop.DevMod {
		return gin.Default()
	}
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
