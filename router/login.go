package router

import (
	"paper-manager/api"
	"paper-manager/middleware"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
}

func (l *LoginRouter) InitLoginRouter(r *gin.RouterGroup) {
	loginApi := new(api.LoginApi)
	r.GET("/getUsers", loginApi.GetUsers)
	r.GET("/getUser", loginApi.GetUserByEmial)
	r.POST("/login", loginApi.Login)
	r.GET("/logout", middleware.JWTAuth(), loginApi.Logout)
}
