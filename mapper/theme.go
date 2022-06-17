package mapper

import (
	"log"

	"github.com/yafeng-Soong/gin-mindmap-manager/database"
	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/theme"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/theme/request"
	"gorm.io/gorm"
)

type ThemeMapper struct{}

func (t *ThemeMapper) SelectById(id int) *theme.Theme {
	theme := &theme.Theme{}
	if err := global.DB.First(theme, id).Error; err != nil {
		// 出错大概率是找不到，也不排除数据库连接出错
		log.Println(err.Error())
		return nil
	}
	return theme
}

func (t *ThemeMapper) UpdateById(updateVo theme.Theme) error {
	tmp := theme.Theme{Name: updateVo.Name, Description: updateVo.Description}
	if err := global.DB.Model(&tmp).Updates(tmp).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *ThemeMapper) ChangeState(theme theme.Theme) error {
	if err := global.DB.Model(&theme).Update("state", theme.State).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *ThemeMapper) SelectPages(p *database.Page[theme.Theme], queryVo request.ThemeQueryVo, userId int) error {
	p.CurrentPage = queryVo.CurrentPage
	p.PageSize = queryVo.PageSize
	query := generateQuery(queryVo, userId)
	err := p.SelectPage(query)
	return err
}

func generateQuery(queryVo request.ThemeQueryVo, userId int) *gorm.DB {
	var state int
	if queryVo.Removed {
		state = theme.REMOVED.Code
	} else {
		state = theme.NORMAL.Code
	}
	query := global.DB.Where(&theme.Theme{State: state, CreatorId: userId}, "state", "creator_id")
	if queryVo.Name != "" {
		query = query.Where("name like ?", "%"+queryVo.Name+"%")
	}
	query.Order("update_time DESC")
	return query
}
