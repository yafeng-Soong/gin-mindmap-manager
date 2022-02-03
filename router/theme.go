package router

import (
	"paper-manager/api"

	"github.com/gin-gonic/gin"
)

type ThemeRouter struct{}

func (t *ThemeRouter) InitRouter(r *gin.Engine) {
	themeApi := new(api.ThemeApi)
	group := r.Group("/theme")
	group.POST("getPageList", themeApi.GetPageList)
}
