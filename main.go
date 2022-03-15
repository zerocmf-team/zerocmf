package main

import (
	"gincmf/app/migrate"
	"gincmf/router"
	cmf "github.com/gincmf/cmf/bootstrap"
)

// test commit
func main() {
	//初始化配置设置
	cmf.Initialize("./conf/config.json")
	//初始化路由设置
	router.ApiListenRouter()
	// 数据库迁移
	migrate.AutoMigrate()
	//启动服务
	cmf.Start()
	//执行数据库迁移
}
