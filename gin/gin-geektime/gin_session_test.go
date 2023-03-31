package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestSession(t *testing.T) {
	r := gin.Default()

	store := cookie.NewStore([]byte("world")) //
	r.Use(sessions.Sessions("Mysession", store))

	r.GET("/abc", func(c *gin.Context) {
		session := sessions.Default(c)
		//panic("test")
		fmt.Println(session.Get("hello")) // world
		//fmt.Println(session.Get("bowen"))
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		} else {
			session.Set("hello", "bowen")
			session.Save()
		}
		c.JSON(200, gin.H{"message": session.Get("hello")})
	})
	r.Run(":8081")
}
