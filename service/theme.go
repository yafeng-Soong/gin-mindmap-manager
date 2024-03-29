package service

import (
	"github.com/yafeng-Soong/gin-mindmap-manager/mapper"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/common/response"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/errors"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/theme"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/theme/request"
	theme_response "github.com/yafeng-Soong/gin-mindmap-manager/model/theme/response"
	user_response "github.com/yafeng-Soong/gin-mindmap-manager/model/user/response"
	paginator "github.com/yafeng-Soong/gorm-paginator"
)

type ThemeService struct{}

var themeMapper mapper.ThemeMapper

func (t *ThemeService) SelectPages(queryVo request.ThemeQueryVo, user user_response.UserInfo) (*response.PageResponse, error) {
	page := paginator.Page[theme.Theme]{}
	err := themeMapper.SelectPages(&page, queryVo, user.Id)
	if err != nil {
		return nil, err
	}
	p := response.NewPageResponse(&page)
	resList := make([]theme_response.ThemeResponse, 0)
	for _, t := range page.Data {
		res := theme_response.NewThemeResponse(t)
		res.Creator = user.Username
		resList = append(resList, *res)
	}
	p.Data = resList
	return p, nil
}

func (t *ThemeService) UpdateTheme(updateVo request.ThemeUpdateVo, userId int) error {
	themeId := updateVo.Id
	if e := themeOperable(themeId, userId); e != nil {
		return e
	}
	name := updateVo.Name
	description := updateVo.Description
	newTheme := &theme.Theme{Id: themeId}
	allBlank := true
	if name != "" {
		newTheme.Name = name
		allBlank = false
	}
	if description != "" {
		newTheme.Description = description
		allBlank = false
	}
	if allBlank {
		err := errors.ERROR
		err.Data = "至少一个更新字段不为空"
		return err
	}
	if e := themeMapper.UpdateById(*newTheme); e != nil {
		err := errors.INNER_ERROR
		return err
	}
	return nil
}

func (t *ThemeService) UpdateThemeState(themeId int, userId int, stateCode int) error {
	if e := themeChangeable(themeId, userId, stateCode); e != nil {
		return e
	}
	tmp := theme.Theme{Id: themeId, State: stateCode}
	if e := themeMapper.ChangeState(tmp); e != nil {
		err := errors.INNER_ERROR
		return err
	}
	return nil
}

func themeOperable(themeId int, userId int) error {
	t := themeMapper.SelectById(themeId)
	err := errors.ERROR
	if t == nil {
		err.Data = "脑图不存在"
		return err
	}
	if t.State != theme.NORMAL.Code {
		err.Data = "脑图当前状态不可操作"
		return err
	}
	if t.CreatorId != userId {
		err.Data = "您没有权限操作该脑图——" + t.Name
		return err
	}
	return nil
}

func themeChangeable(themeId int, userId int, state int) error {
	t := themeMapper.SelectById(themeId)
	err := errors.ERROR
	if t == nil {
		err.Data = "脑图不存在"
		return err
	}
	if t.State == state {
		// 当前状态与目标状态一致
		err.Data = "脑图状态有误"
		return err
	}
	if t.CreatorId != userId {
		err.Data = "您没有权限操作该脑图——" + t.Name
		return err
	}
	return nil
}
