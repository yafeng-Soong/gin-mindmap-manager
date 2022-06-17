package initial

import (
	"fmt"

	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.DBname,
		global.CONFIG.Mysql.Charset,
	)
	// 使用第三方logger将gorm的log写入到zap
	logger := zapgorm2.New(global.LOG)
	logger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.LogMode(gormlogger.Info),
	})
	if err != nil {
		global.LOG.Fatal("数据库连接错误！", zap.Error(err))
	}
	global.DB = db
}
