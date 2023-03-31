package demo

import (
	"github.com/gin-gonic/gin"
	"testing"
)

// @program:   go-web-exercises
// @file:      demo1_test.go
// @author:    bowen
// @time:      2023-03-30 17:39
// @description:
func TestGindemo1(t *testing.T) {
	// 创建默认的路由引擎
	r := gin.Default()
	// GET: 请求方式，/hello请求路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数，之前学过的匿名函数的用法
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,Gin",
		})
	})
	//启动，默认是本地8080启动
	r.Run()
}

func TestGinDemo2(t *testing.T) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello, gin")
	})
	r.Run()
}
