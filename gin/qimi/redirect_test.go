package qimi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	//创建一个默认的路由引擎
	r := gin.Default()
	// GET: 请求方式，/hello请求路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务

	r.GET("/abc", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		//r.HandleContext(c)
	})
	r.Run()
}

func TestRedirect2(t *testing.T) {
	//创建一个默认的路由引擎
	r := gin.Default()
	// GET: 请求方式，/hello请求路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务

	r.GET("/abc", func(c *gin.Context) {
		//c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})
	r.Run()
}
