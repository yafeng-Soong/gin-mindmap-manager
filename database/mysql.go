package database

import (
	// 导入gorm工具包

	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Close() {
	db, _ := DB.DB()
	db.Close()
}

func InitMysql() error {
	dsn := "root:fuck@you@tcp(127.0.0.1:3306)/paper_manager_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}
	DB = db
	return nil
}

func SelectPage(page *Page, query *gorm.DB, model interface{}) (e error) {
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
