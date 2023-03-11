package handle

import (
	"log"
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

		log.Println(c.Request.Header.Get(consts.Authorization))
		log.Println(c.Request.Header.Get(consts.Client))
		var token, clientID string
		if token, clientID = c.Request.Header.Get(consts.Authorization),
			c.Request.Header.Get(consts.Client); token == "" || clientID == "" {
			c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "获取 JWT 或者 clientID 失败", nil))
			return
		}

		uid, err := com.jwt.Parse(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, NewApiResult(501, "JWT 解析失败", nil))
			return
		}

		c.Request.Header.Set("userID", uid)
		c.Request.Header.Set("clientID", clientID)

		//user, err := com.user.GetByID(c, uid)
		//if err != nil {
		//	c.AbortWithStatusJSON(200, gin.H{
		//		"msg": "请登陆后重试",
		//	})
		//	return
		//}
	}
}
