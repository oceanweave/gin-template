package main

import (
	"fmt"
	"gin-template/conf"
	"gin-template/internal/dao/db"
	"gin-template/pkg/httpserver"
	"gin-template/pkg/logger"
	"gin-template/routes"
)

func main() {
	// 初始化配置
	if err := conf.LoadConf(); err != nil {
		panic(fmt.Errorf("load conf failed, reason:%s", err.Error()))
	}

	// 初始化日志
	if err := logger.InitLogger(); err != nil {
		panic(fmt.Errorf("init logger error, reason:%v", err))
	}

	// 初始化数据库
	if err := db.InitMysql(); err != nil {
		panic(fmt.Errorf("init mysql failed, reason:%s", err.Error()))
	}

	// 创建gin路由
	engine := routes.NewGinRouter()
	// 注册路由
	routes.Register(engine)

	// 创建http server
	server := httpserver.NewServer(conf.ServerConfig.Host, conf.ServerConfig.Port, engine)

	// 启动server
	httpserver.ListenAndServe(server)

	// 等待服务退出
	httpserver.WaitForShutdown(server, func() {
		db.CloseMysql()
		fmt.Println("close mysql connection.")
	})
}
