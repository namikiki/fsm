package middlewares

import (
	"fsm/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyUserToken() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//
	//	var token string
	//	if token = c.GetHeader(consts.Authorization); token == "" ||
	//		c.GetHeader(consts.Client) == "" {
	//		c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "获取 JWT 或者 clientID 失败", nil))
	//		return
	//	}
	//
	//	_, err := utils.ValidateToken((token)
	//	if err != nil {
	//		c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "JWT 解析失败", nil))
	//		return
	//	}
	//
	//	c.Request.Header.Set("userID", uid)
	//}
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		_, err := utils.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}

}
