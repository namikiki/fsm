package handle

import (
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

		//token := c.Request.Header.Get(consts.Authorization)
		//if token == "" {
		//	c.AbortWithStatusJSON(200, gin.H{
		//		"msg": "请登陆后重试",
		//	})
		//	return
		//}
		//
		//uid, err := com.jwt.Parse(c, token)
		//if err != nil {
		//	c.AbortWithStatusJSON(200, gin.H{
		//		"msg": "请登陆后重试",
		//	})
		//	return
		//}
		//
		//user, err := com.user.GetByID(c, uid)
		//if err != nil {
		//	c.AbortWithStatusJSON(200, gin.H{
		//		"msg": "请登陆后重试",
		//	})
		//	return
		//}
		//
		//c.Set("user", user)
	}
}
