package qimi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestRoute(t *testing.T) {
	r := gin.Default()
	r.GET("/var", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "这里是变量",
		})
	})

	r.Any("/anymenthod", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "这里返回any页面",
		})
	})

	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "没有找到对应页面",
	//	})
	//})

	r.NoRoute(func(c *gin.Context) {
		// 这个方法倒是挺有意思的，直接替换掉了
		c.String(http.StatusNotFound, "没有该页面")
	})
	r.Run()
}

func TestRouteGroup(t *testing.T) {
	r := gin.Default()
	userGroup := r.Group("/user")
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
	{
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
	}
	r.Run()
}
