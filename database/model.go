package database

// import (
// 	"gorm.io/gorm"
// )

// type Page[T any] struct {
// 	CurrentPage int64
// 	PageSize    int64
// 	Total       int64
// 	Pages       int64
// 	Data        []T
// }

// func (page *Page[T]) SelectPage(query *gorm.DB) (e error) {
// 	var model T
// 	query.Model(&model).Count(&page.Total)
// 	if page.Total == 0 {
// 		page.Data = []T{}
// 		return
// 	}
// 	e = query.Model(&model).Scopes(Paginate(page)).Find(&page.Data).Error
// 	return
// }

// func Paginate[T any](page *Page[T]) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if page.CurrentPage <= 0 {
// 			page.CurrentPage = 0
// 		}
// 		switch {
// 		case page.PageSize > 100:
// 			page.PageSize = 100
// 		case page.PageSize <= 0:
// 			page.PageSize = 10
// 		}
// 		page.Pages = page.Total / page.PageSize
// 		if page.Total%page.PageSize != 0 {
// 			page.Pages++
// 		}
// 		p := page.CurrentPage
// 		if page.CurrentPage > page.Pages {
// 			p = page.Pages
// 		}
// 		size := page.PageSize
// 		offset := int((p - 1) * size)
// 		return db.Offset(offset).Limit(int(size))
// 	}
// }
