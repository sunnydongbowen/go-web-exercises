package demo

import (
	"github.com/gin-gonic/gin"
	"testing"
)

// @program:   go-web-exercises
// @file:      demo4_test.go
// @author:    bowen
// @time:      2023-03-31 9:33
// @description: 路由组

func TestRouteGroup(t *testing.T) {
	r := gin.Default()
	userGroup := r.Group("/user")
	// 不用{}也可
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "这里user的index页面",
			})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "这里是user的login页面",
			})
		})
	}

	shopGroup := r.Group("/shop")
	shopGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mesage": "这是shop的index页面",
		})
	})
	shopGroup.GET("/cart", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "这是shop的cart页面",
		})
	})
	r.Run()
}
