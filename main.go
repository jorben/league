package main

import (
	"github.com/gin-gonic/gin"
	"league/config"
	"league/database"
	"league/log"
	"league/middleware"
	"league/model"
	"league/router"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	// 初始化日志
	log.InitLogger(cfg.Log)

	// 连接DB
	err = database.InitMysql(cfg.Db)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// 同步schema
	if err := model.AutoMigrate(); err != nil {
		panic(err)
	}

	// 构建Gin实例
	s := gin.New()
	s.Use(middleware.RequestId(), middleware.Logger(), middleware.Auth(), gin.Recovery())

	// 注册路由
	router.SetupRouter(s)

	// 启动服务
	if err := s.Run(":8080"); err != nil {
		panic(err)
		return
	}
}
