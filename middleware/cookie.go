package middleware

import (
	"log"
	"paper-manager/model/common/response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Println("进入了Session中间件")
		session := sessions.Default(c)
		// log.Println(session)
		if session.Get("currentUser") == nil {
			res := response.UNAUTHORIZED()
			response.FailWithMessage(res.Code, res.Msg, c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func SetSession(data interface{}, c *gin.Context) {
	session := sessions.Default(c)
	session.Set("currentUser", data)
	session.Options(sessions.Options{MaxAge: 259200})
	if err := session.Save(); err != nil {
		log.Println(err)
	}

}

func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
