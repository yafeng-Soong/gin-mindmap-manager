package api

import (
	"log"
	"paper-manager/model/common/response"
	"paper-manager/model/errors"
	"paper-manager/model/theme"
	"paper-manager/model/theme/request"
	"paper-manager/service"
	"paper-manager/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThemeApi struct{}

var themeService service.ThemeService

func (t *ThemeApi) GetPageList(c *gin.Context) {
	var queryVo request.ThemeQueryVo
	if e := c.ShouldBindJSON(&queryVo); e != nil {
		log.Println(e.Error())
		c.Error(errors.GetError(errors.VALID_ERROR, e.Error()))
		return
	}
	currentUser, _ := utils.GetCurrentUser(c)
	// pageResponse, err := themeService.SelectPageList(queryVo, *currentUser)
	pageResponse, err := themeService.SelectPages(queryVo, *currentUser)
	if err != nil {
		c.Error(err)
		return
	}
	response.OkWithData(pageResponse, c)
}

func (t *ThemeApi) UpdateTheme(c *gin.Context) {
	var updateVo request.ThemeUpdateVo
	if e := c.ShouldBindJSON(&updateVo); e != nil {
		log.Println(e.Error())
		c.Error(errors.GetError(errors.VALID_ERROR, e.Error()))
		return
	}
	currentUser, _ := utils.GetCurrentUser(c)
	if err := themeService.UpdateTheme(updateVo, currentUser.Id); err != nil {
		c.Error(err)
		return
	}
	response.OkWithData("更新脑图信息成功", c)
}

func (t *ThemeApi) RemoveTheme(c *gin.Context) {
	var themeId int
	themeId, e := strconv.Atoi(c.Query("themeId"))
	if e != nil {
		log.Println(e.Error())
		c.Error(errors.GetError(errors.VALID_ERROR, e.Error()))
		return
	}
	currentUser, _ := utils.GetCurrentUser(c)
	removeCode := theme.REMOVED.Code
	if err := themeService.UpdateThemeState(themeId, currentUser.Id, removeCode); err != nil {
		c.Error(err)
		return
	}
	response.OkWithData("删除成功", c)
}

func (t *ThemeApi) RecoverTheme(c *gin.Context) {
	var themeId int
	themeId, e := strconv.Atoi(c.Query("themeId"))
	if e != nil {
		log.Println(e.Error())
		c.Error(errors.GetError(errors.VALID_ERROR, e.Error()))
		return
	}
	currentUser, _ := utils.GetCurrentUser(c)
	recoverCode := theme.NORMAL.Code
	if err := themeService.UpdateThemeState(themeId, currentUser.Id, recoverCode); err != nil {
		c.Error(err)
		return
	}
	response.OkWithData("恢复成功", c)
}
