package qimi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestValidate(t *testing.T) {
	r := gin.Default()
	r.POST("/api/test", testHandler)

	r.Run()

}

type Param struct {
	ID      int     `json:"-"`
	Name    *string `json:"name" binding:"required"` // required表示必须传，否则会报错
	Age     *int    `json:"age" binding:"required"`
	Married *bool   `json:"married" binding:"required"`
}

func testHandler(c *gin.Context) {
	// 参数获取与参数校验
	var p Param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, "参数错误:"+err.Error())
		return
	}
	fmt.Printf("%#v\n", p)
	// 获取有效参数
	c.JSON(http.StatusOK, p)
}
