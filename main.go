package main

import (
	"fmt"

	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"github.com/yafeng-Soong/gin-mindmap-manager/initial"
	"github.com/yafeng-Soong/gin-mindmap-manager/router"
)

func main() {
	initial.InitialEnv() // 初始化所有全局参数
	db, _ := global.DB.DB()
	defer db.Close()
	r := router.SetupRouter()
	addr := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	r.Run(addr)
}
