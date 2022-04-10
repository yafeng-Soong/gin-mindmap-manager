package theme

import (
	"log"
	"paper-manager/database"
	"paper-manager/model/theme/request"
	"time"

	"gorm.io/gorm"
)

type Theme struct {
	Id          int
	CreatorId   int
	Name        string
	Description string
	State       int
	CreateTime  time.Time
	UpdateTime  time.Time
}

type ThemeState struct {
	State string
	Code  int
}

var (
	NORMAL  = &ThemeState{State: "正常", Code: 0}
	REMOVED = &ThemeState{State: "被删除", Code: 1}
)

func (t *Theme) TableName() string {
	return "theme"
}

func (t *Theme) SelectById(id int) *Theme {
	theme := &Theme{}
	if err := database.DB.First(theme, id).Error; err != nil {
		// 出错大概率是找不到，也不排除数据库连接出错
		log.Println(err.Error())
		return nil
	}
	return theme
}

func (t *Theme) UpdateById(theme Theme) error {
	tmp := Theme{Name: theme.Name, Description: theme.Description}
	if err := database.DB.Model(&theme).Updates(tmp).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *Theme) ChangeState(theme Theme) error {
	if err := database.DB.Model(&theme).Update("state", theme.State).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (t *Theme) SelectPages(p *database.Page[Theme], queryVo request.ThemeQueryVo, userId int) error {
	p.CurrentPage = queryVo.CurrentPage
	p.PageSize = queryVo.PageSize
	query := generateQuery(queryVo, userId)
	err := p.SelectPage(query)
	return err
}

func generateQuery(queryVo request.ThemeQueryVo, userId int) *gorm.DB {
	var state int
	if queryVo.Removed {
		state = REMOVED.Code
	} else {
		state = NORMAL.Code
	}
	query := database.DB.Where(&Theme{State: state, CreatorId: userId}, "state", "creator_id")
	if queryVo.Name != "" {
		query = query.Where("name like ?", "%"+queryVo.Name+"%")
	}
	query.Order("update_time DESC")
	return query
}
