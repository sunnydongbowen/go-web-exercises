package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func TestGlobal(t *testing.T) {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StatCost())
	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}

func TestLocal(t *testing.T) {
	r := gin.Default()
	r.GET("/test", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/xyz", StatCost(), func(c *gin.Context) {
		//name := c.MustGet("name").(string)
		//log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "hello xyz",
		})
	})

	r.Run()

}

//func TestGroup(t *testing.T) {
//
//}
