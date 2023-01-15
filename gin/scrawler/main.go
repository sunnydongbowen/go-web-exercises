package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @program:     go-web-exercises
// @file:        main.go
// @author:      bowen
// @create:      2023-01-08 22:12
// @description: 写的爬虫用的

func main() {
	r := gin.Default()
	//初始化连接

	r.LoadHTMLFiles("gin/scrawler/test.html")

	r.GET("/main", func(c *gin.Context) {
		c.String(200, "I am in main")

	})

	r.GET("/bobo", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.HTML(http.StatusOK, "test.html", nil)
	})

	r.GET("/tom", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.HTML(http.StatusOK, "test.html", nil)
	})

	r.GET("/jay", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.HTML(http.StatusOK, "test.html", nil)
	})
	//启动服务
	r.Run()
}
