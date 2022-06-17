package initial

func InitialEnv() {
	InitViper() // 最先初始化Viper
	InitZap()   // 第二步初始化全局log
	InitMysql() // 初始化Mysql连接
}
