package api

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//return gin.New()
	return gin.Default()
}

//
//func Start(app *gin.Engine) error {
//	return app.Run(":8080")
//}
