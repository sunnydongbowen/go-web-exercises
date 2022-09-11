package gin_geektime

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestContoller_GetUser(t *testing.T) {
	g := gin.Default()
	ctrl := &UserControlller{}
	// 为get请求时走到这里，其实这里本质上和下面的post没有什么区别
	// 想说，MVC 模式，应该是用户在使用 Web 框架的时候，组织自己代码使用的设计模式
	// 而不是我们框架应该内置的模式。因为有些用户喜欢 MVC，有些不喜欢，框架不应该有 MVC 的假设
	g.GET("/user/", ctrl.GetUser)
	// 为post请求时走到这里
	g.POST("/user/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello %s", "bowen")
	})

	g.GET("/static", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello get")
		// 读文件
		// 写响应
	})

	_ = g.Run(":8082")
}
