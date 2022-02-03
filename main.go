package main

import (
	"log"
	"paper-manager/database"
	"paper-manager/router"
)

func main() {
	if err := database.InitMysql(); err != nil {
		log.Fatalln("数据库连接出错")
	}
	defer database.Close()
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
