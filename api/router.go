package api

import (
	"log"
	"net/http"

	"fsm/api/handle"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AddRouters(app *gin.Engine, user handle.User, file handle.File, dir handle.Dir, syncTask handle.SyncTask, common handle.Common) {
	store := cookie.NewStore([]byte("secret"))

	app.Use(common.VerifyUserToken())
	
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	// test route
	{
		app.GET("/test", func(c *gin.Context) {
			code := c.Query("testcode")
			if code == "123" {
				c.AbortWithStatusJSON(http.StatusOK, "success")

			} else {
				c.AbortWithStatusJSON(http.StatusOK, "fail")
			}
			log.Println("不应该执行")
		})

	}

	// file
	{
		app.GET("/file/open/:fileID", file.Open)
		app.GET("/file/getMetadata", file.GetMetadata)
		app.POST("/file/create", file.Create)
		app.GET("/file/get/all/bySyncID", file.GetAllFileBySyncID)
	}

	// dir
	{
		app.POST("/dir/create", dir.Create)
		app.GET("/dir/delete", dir.Delete)
		app.GET("/dir/read", dir.ReadDir)
		app.GET("/dir/getAllDirByPath", dir.GetAllDirByPath)
		app.GET("/dir/getAllDirBySyncID", dir.GetAllDirBySyncID)
	}

	//syncTask
	{
		app.POST("/synctask/create", syncTask.Create)
		app.GET("/synctask/delete", syncTask.Delete)
		app.GET("/synctask/get", syncTask.Get)
		app.GET("/synctask/getAll", syncTask.GetAll)
	}

	app.Use(sessions.Sessions("sessionId", store))
	app.GET("/websocketconn", user.WebsocketConn)

	app.POST("/login", user.Login)
	app.POST("/register", user.Register)
	app.GET("/delete", user.Delete)

	app.POST("/filestore", file.Create)
}
