package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// @program:   go-web-exercises
// @file:      demo3_test.go
// @author:    bowen
// @time:      2023-03-30 21:33
// @description:route

func TestRoute(t *testing.T) {
	r := gin.Default()
	r.GET("/var", func(c *gin.Context) {
		c.JSON(200, "你好") // 只返回你好
	})

	r.Any("/AnyMethod", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "这里返回any页面",
		})
	})

	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "没有找到对应页面",
	//	})
	//
	//})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "没有该页面")
	})

	r.Run()

}
