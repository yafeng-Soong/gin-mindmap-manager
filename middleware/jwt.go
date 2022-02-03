package middleware

import (
	"paper-manager/model/common/response"
	"paper-manager/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		res := response.UNAUTHORIZED()
		if len(auth) == 0 {
			c.Abort()
			response.FailWithMessage(res.Code, res.Msg, c)
			return
		}
		// auth = strings.Fields(auth)[1]
		// 校验token
		claims, err := utils.ParseToken(auth)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				newToken, _ := utils.RenewToken(claims)
				if newToken != "" {
					c.Header("newtoken", newToken)
					c.Request.Header.Set("Authorization", newToken)
					c.Next()
					return
				}
			}
			c.Abort()
			message := err.Error()
			response.FailWithMessage(res.Code, message, c)
			return
		}
		c.Next()
	}
}
