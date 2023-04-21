package handle

import (
	"net/http"

	"fsm/api/consts"
	"fsm/pkg/domain"
	"fsm/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Common struct {
	jwt  jwt.Service
	user domain.UserRepository
}

func NewCommon(jwt jwt.Service, user domain.UserRepository) Common {
	return Common{jwt: jwt, user: user}
}

func (com *Common) VerifyUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		var token string
		if token = c.GetHeader(consts.Authorization); token == "" ||
			c.GetHeader(consts.Client) == "" {
			c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "获取 JWT 或者 clientID 失败", nil))
			return
		}

		uid, err := com.jwt.Parse(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "JWT 解析失败", nil))
			return
		}

		c.Request.Header.Set("userID", uid)
	}
}
