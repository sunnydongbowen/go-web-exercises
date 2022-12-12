package router

import (
	"github.com/gin-gonic/gin"
	controller2 "go-web-exercises/gin_zap/controller"
	"go-web-exercises/gin_zap/middleware"
	"go-web-exercises/gin_zap/setting"
	"go.uber.org/zap"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.Conf.Mode)

	r := gin.New()
	// 注册两个自定义的中间件
	r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), false))

	r.GET("/hello", controller2.HelloHandler)
	r.GET("/ping", controller2.PingHandler)
	r.GET("/login", controller2.LoginHandler)

	return r
}
