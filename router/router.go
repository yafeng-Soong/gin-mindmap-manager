package router

import (
	// "encoding/gob"
	"github.com/yafeng-Soong/gin-mindmap-manager/middleware"

	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)
var (
	loginRouter LoginRouter
	testRouter  TestRouter
	themeRouter ThemeRouter
)

func SetupRouter() *gin.Engine {
	// gob.Register(user_response.UserInfo{})
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.Cors())
	// store := cookie.NewStore([]byte("snaosnca"))
	// r.Use(sessions.Sessions("SESSIONID", store))
	// r.Use(middleware.Cookie())
	publicGroup := r.Group("")
	loginRouter.InitLoginRouter(publicGroup)
	r.Use(middleware.JWTAuth())
	// loginRouter.InitLoginRouter(r)
	testRouter.InitTestRouter(r)
	themeRouter.InitRouter(r)

	return r
}
