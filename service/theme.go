package service

import (
	"paper-manager/model/common/response"
	"paper-manager/model/theme"
	"paper-manager/model/theme/request"
	theme_response "paper-manager/model/theme/response"
	user_response "paper-manager/model/user/response"
)

type ThemeService struct{}

var themeModel theme.Theme

func (t *ThemeService) SelectPageList(queryVo request.ThemeQueryVo, user user_response.UserInfo) (*response.PageResponse, error) {
	resList := make([]theme_response.ThemeResponse, 0)
	p := response.NewPageResponse(queryVo.CurrentPage, queryVo.PageSize)
	p.Total = themeModel.CountAll(queryVo, user.Id)
	if !p.CountPages() {
		p.Data = resList
		return p, nil
	}
	themeList, err := themeModel.SelectList(queryVo, user.Id)
	// err := themeModel.SelectPageList(p, queryVo, user.Id)
	if err != nil {
		return nil, err
	}
	for _, theme := range themeList {
		res := theme_response.NewThemeResponse(theme)
		res.Creator = user.Username
		resList = append(resList, *res)
	}
	p.Data = resList
	return p, nil
}
