package api

import (
	"log"
	"paper-manager/model/common/response"
	"paper-manager/model/theme/request"
	"paper-manager/service"
	"paper-manager/utils"

	"github.com/gin-gonic/gin"
)

type ThemeApi struct{}

var themeService service.ThemeService

func (t *ThemeApi) GetPageList(c *gin.Context) {
	var queryVo request.ThemeQueryVo
	var res *response.ResponseEnums
	if e := c.ShouldBindJSON(&queryVo); e != nil {
		log.Println(e.Error())
		res = response.VALID_ERROR()
		response.FailWithMessage(res.Code, res.Msg, c)
		return
	}
	currentUser, _ := utils.GetCurrentUser(c)
	pageResponse, err := themeService.SelectPageList(queryVo, *currentUser)
	if err != nil {
		c.Error(err)
		return
	}
	response.OkWithData(pageResponse, c)
}
