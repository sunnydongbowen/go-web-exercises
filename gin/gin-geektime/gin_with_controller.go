package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControlller struct {
}

func (c *UserControlller) GetUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello,world")
}
