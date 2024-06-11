package middlewares

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Authorization = "authorization"
)

type Auth struct {
	jwt *services.JWTService
}

func NewAuth(jwt *services.JWTService) *Auth {
	return &Auth{jwt: jwt}
}

func (auth *Auth) VerifyUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		var token string
		if token = c.GetHeader(Authorization); token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "error authorization",
			})
			return
		}

		uid, err := auth.jwt.Parse(c, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "jwt 解析失败",
			})
			return
		}

		c.Request.Header.Set("userID", uid)
	}
}
