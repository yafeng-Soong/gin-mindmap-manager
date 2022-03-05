package database

import (
	"reflect"

	"gorm.io/gorm"
)

type Page struct {
	CurrentPage int64
	PageSize    int64
	Total       int64
	Pages       int64
	Data        []interface{}
}

func (page *Page) SelectPage(query *gorm.DB, model interface{}) (e error) {
	e = nil
	query.Model(&model).Count(&page.Total)
	if page.Total == 0 {
		page.Data = []interface{}{}
		return
	}
	t := reflect.TypeOf(model)
	list := reflect.Zero(reflect.SliceOf(t)).Interface()
	e = query.Model(&model).Scopes(Paginate(page)).Find(&list).Error
	if e != nil {
		return
	}
	// log.Println(list)
	page.Data = toSlice(list)
	return
}

func toSlice(arr interface{}) []interface{} {
	ret := make([]interface{}, 0)
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		ret = append(ret, arr)
		return ret
	}
	l := v.Len()
	for i := 0; i < l; i++ {
		ret = append(ret, v.Index(i).Interface())
	}
	return ret
}

func Paginate(page *Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.CurrentPage <= 0 {
			page.CurrentPage = 0
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.CurrentPage
		if page.CurrentPage > page.Pages {
			p = page.Pages
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
