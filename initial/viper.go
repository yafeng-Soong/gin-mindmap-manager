package initial

import (
	"log"

	"github.com/spf13/viper"
	"github.com/yafeng-Soong/gin-mindmap-manager/global"
)

func InitViper() {
	vi := viper.New()
	vi.SetConfigName("config")
	vi.SetConfigType("yaml")
	vi.AddConfigPath("conf")
	if err := vi.ReadInConfig(); err != nil {
		log.Println(err.Error())
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Fatal("未找到配置文件！", err.Error())
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Fatal("读取配置文件出错！", err.Error())
		}
	}
	if err := vi.Unmarshal(&global.CONFIG); err != nil {
		log.Println(err.Error())
		log.Fatal("解析配置文件出错！")
	}
	global.VIPER = vi
}
