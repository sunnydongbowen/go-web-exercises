package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

// @program:   go-web-exercises
// @file:      demo2_test.go
// @author:    bowen
// @time:      2023-03-30 20:24
// @description:restfulAPI

func TestRestful(t *testing.T) {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run()
}

func TestH(t *testing.T) {
	m := map[string]any{"message": "GET", "Code": 200}
	fmt.Println(m)
}
