package gin_geektime

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGinSession(t *testing.T) {
	r := gin.Default()
	store := cookie.NewStore([]byte("world")) // cookie里面带了session
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		} else {
			session.Set("hello", "bowen")
			session.Save()
		}
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8080")
}
