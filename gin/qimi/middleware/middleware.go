package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/xyz" {
			c.Next()
			//c.Abort()
			return
			// abort()终端当前的handlerfunc
		}
		start := time.Now()
		// 闭包的用法，内层函数使用外层函数的变量 format string,可以传参数
		//log.Println(start.Format(format))

		c.Set("name", "bowen")
		// 调用该请求的剩余处理程序
		c.Next()
		//计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
