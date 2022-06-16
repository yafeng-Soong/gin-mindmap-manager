package router

import (
	"github.com/yafeng-Soong/gin-mindmap-manager/api"

	"github.com/gin-gonic/gin"
)

type ThemeRouter struct{}

func (t *ThemeRouter) InitRouter(r *gin.Engine) {
	themeApi := new(api.ThemeApi)
	group := r.Group("/theme")
	group.POST("/getPageList", themeApi.GetPageList)
	group.POST("/update", themeApi.UpdateTheme)
	group.DELETE("/remove", themeApi.RemoveTheme)
	group.GET("/recover", themeApi.RecoverTheme)
}
