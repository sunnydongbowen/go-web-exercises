package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

// @program:     go-web-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-12-14 16:47
// @description: 中间件记录返回响应体，写入数据库
type param struct {
	X int `json:"x" binding:"required"`
	Y int `json:"y" binding:"required"`
}

type bodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架responseWriter
	body               *bytes.Buffer // 我们记录用的response
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 我们记录一一份
	return w.ResponseWriter.Write(b) //真正写入响应数据
}

func recoredResponseMilldleware(c *gin.Context) {
	fmt.Println("in recoredResponseMilldleware")
	// c.Writer 换成自定义的
	newWriter := &bodyLogWriter{
		body:           bytes.NewBuffer([]byte{}),
		ResponseWriter: c.Writer,
	}
	c.Writer = newWriter
	c.Next()
	// 记录addHandler返回的响应数据是什么
	// 1.获取返回响应数据
	// c.Request.Response.Body) 响应体， io.ReadCloser接口类型
	//b, _ := io.ReadAll(c.Request.Response.Body)
	// 满足条件才记录
	statusCode := c.Writer.Status()
	//if c.Request.Response.StatusCode == 400 {
	if statusCode == 400 {
		// gin框架需要把这个响应数据通过网络返回给请求客户端
		// 在这里把应  该返回给浏览器的响应给读完了，读到io.EOF
		//fmt.Println("xxx")
		b, _ := io.ReadAll(newWriter.body)
		fmt.Printf("--。>%s\n", b)
	}

	fmt.Println("out recoredResponseMilldleware")
}

func addHandler(c *gin.Context) {
	fmt.Println("in addHandler")
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数有误",
			"err":  err.Error(),
		})
		return
	}
	res := p.X + p.Y
	c.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

func main() {
	r := gin.Default()

	r.POST("/add", recoredResponseMilldleware, addHandler)
	r.Run()

}
