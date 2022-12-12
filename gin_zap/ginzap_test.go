package qimi

import (
	"fmt"
	"go-web-exercises/gin_zap/dao/mysql"
	"go-web-exercises/gin_zap/logger"
	"go-web-exercises/gin_zap/router"
	"go-web-exercises/gin_zap/setting"
	"go.uber.org/zap"
	"testing"
)

// @program:     go-web-exercises
// @file:        ginzap_test.go
// @author:      bowen
// @create:      2022-11-12 11:28
// @description: 七米视频，gin和zap,项目通用开发框架，涉及gin，viper，数据库，zap日志等，是一个大的架子。

func TestZap(t *testing.T) {

	// 1. 加载配置,这里我传了绝对路径。
	// 这里直接从配置文件读取了，写死了。
	err := setting.Init("E:\\go-web-exercises\\gin\\qimi\\gin_zap\\conf\\config.yaml")
	//if len(os.Args[1]) < 2 {
	//	panic("程序执行时必须通过命令行指定配置文件")
	//}
	//err := setting.Init(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Conf.Name)

	//2. 初始化日志模块，注意命名
	err = logger.Init()
	if err != nil {
		panic(err)
	}

	// 3. 数据库初始化
	err = mysql.Init()
	if err != nil {
		zap.L().Error("mysql.Init failed", zap.Error(err))
	}

	// 4. 路由初始化
	r := router.Setup()
	// 5. 程序启动

	////gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//// Ginlogger 是为了将gin的日志也写到日志文件里, Ginrecovery是为了panic后恢复程序
	//r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), false))
	//r.GET("/hello", func(c *gin.Context) {
	//	zap.L().Info("req hello")
	//	panic("enene")
	//})
	//r.GET("/bowen", func(c *gin.Context) {
	//	c.String(200, "博文你好")
	//})

	r.Run()
}
