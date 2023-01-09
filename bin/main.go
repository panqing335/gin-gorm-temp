package main

import (
	Cron "temp/app/common/cron"
	Grpc "temp/app/common/grpc"
	Mysql "temp/app/common/mysql"
	Redis "temp/app/common/redis"
	Router "temp/app/router"
	Config "temp/config"
)

func main() {
	Config.Setup() // 初始化配置
	Redis.Setup()  // 初始化redis
	Mysql.Setup()  // 初始化DB
	Grpc.Setup()   // 初始化Grpc
	Cron.Setup()   // 初始化定时任务
	Router.Setup() // 路由，开启服务
}
