package api

import (
	"log"
	"net/http"

	"fsm/api/handle"
	"fsm/pkg/types"

	"github.com/gin-gonic/gin"
)

func AddRouters(router *gin.Engine, user handle.User, file handle.File, dir handle.Dir, syncTask handle.SyncTask, common handle.Common) {
	//store := cookie.NewStore([]byte("secret"))
	//app.Use(sessions.Sessions("sessionId", store))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	app := router.Group("")
	app.Use(common.VerifyUserToken())

	{ // test route
		app.GET("/test", func(c *gin.Context) {

			log.Println("auth", c.Request.Header.Get("authorization"))
			log.Println("userid", c.Request.Header.Get("userID"))
			log.Println("clientID", c.Request.Header.Get("clientID"))

			code := c.Query("testcode")
			if code == "123" {
				c.AbortWithStatusJSON(http.StatusOK, "success")

			} else {
				c.AbortWithStatusJSON(http.StatusOK, "fail")
			}
			log.Println("不应该执行")
		})
	}

	{ // file
		app.GET("/file/open/:fileID", file.Open)
		app.GET("/file/getMetadata", file.GetMetadata)
		app.POST("/file/create", file.Create)
		app.GET("/file/get/all/bySyncID/:syncID", file.GetAllFileBySyncID)
	}

	{ // dir
		app.POST("/dir/create", dir.Create)
		app.GET("/dir/delete", dir.Delete)
		app.GET("/dir/read", dir.ReadDir)
		app.GET("/dir/getAllDirByPath", dir.GetAllDirByPath)
		app.GET("/dir/getAllDirBySyncID/:syncID", dir.GetAllDirBySyncID)
	}

	{ //syncTask
		app.POST("/synctask/create", syncTask.Create)
		app.GET("/synctask/delete/:syncID", syncTask.Delete)
		app.GET("/synctask/get/:syncID", syncTask.Get)
		app.GET("/synctask/getAll", syncTask.GetAll)
	}

	router.GET("/websocket/connect/:user/:client", user.WebsocketConn)

	router.POST("/login", user.Login)
	router.POST("/register", user.Register)
	app.GET("/delete", user.Delete)

	app.POST("/filestore", file.Create)
}

func New(conf *types.Config) *gin.Engine {
	if conf.Develop.DevMod {
		return gin.Default()
	}
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
