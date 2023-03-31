package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// @program:   go-web-exercises
// @file:      demo5_test.go
// @author:    bowen
// @time:      2023-03-31 11:21
// @description: redirect

func TestRedirect(t *testing.T) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})

	r.GET("/abc", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.Run()
}

func TestRedi2(t *testing.T) {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})

	r.GET("/abc", func(c *gin.Context) {
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})
	r.Run()
}
