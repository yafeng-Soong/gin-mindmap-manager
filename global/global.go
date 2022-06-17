package global

import (
	"github.com/spf13/viper"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	VIPER     *viper.Viper   // 全局viper，留做热更新
	DB        *gorm.DB       // 数据库对象
	CONFIG    *config.Config // 所有配置项
	LOG       *zap.Logger
	SUGAR_LOG *zap.SugaredLogger
)
