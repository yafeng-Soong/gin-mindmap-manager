package theme

import (
	"log"
	"paper-manager/database"
	"paper-manager/model/common/response"
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

func (t *Theme) TableName() string {
	return "theme"
}

func (t *Theme) CountAll(queryVo request.ThemeQueryVo, userId int) int64 {
	var total int64
	query := generateQuery(queryVo, userId)
	query.Model(&Theme{}).Count(&total)
	return total
}

func (t *Theme) SelectList(queryVo request.ThemeQueryVo, userId int) ([]Theme, error) {
	themeList := []Theme{}
	query := generateQuery(queryVo, userId)
	err := query.Order("update_time desc").Limit(int(queryVo.PageSize)).Offset(int((queryVo.CurrentPage - 1) * queryVo.PageSize)).Find(&themeList).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return themeList, nil
}

func (t *Theme) SelectPageList(p *response.PageResponse, queryVo request.ThemeQueryVo, userId int) error {
	themeList := &[]Theme{}
	query := generateQuery(queryVo, userId)
	query.Model(&Theme{}).Count(&p.Total)
	if !p.CountPages() {
		p.Data = themeList
		return nil
	}
	err := query.Order("update_time desc").Limit(int(queryVo.PageSize)).Offset(int((queryVo.CurrentPage - 1) * queryVo.PageSize)).Find(&themeList).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	p.Data = themeList
	return nil
}

func generateQuery(queryVo request.ThemeQueryVo, userId int) *gorm.DB {
	var state int
	if queryVo.Removed {
		state = 1
	} else {
		state = 0
	}
	query := database.DB.Where(&Theme{State: state, CreatorId: userId})
	if queryVo.Name != "" {
		query = query.Where("name like ?", "%"+queryVo.Name+"%")
	}
	return query
}
