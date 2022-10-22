package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func TestUploadMulti(t *testing.T) {
	router := gin.Default()

	// 加载模板
	router.LoadHTMLFiles("upload.html")
	// 上传文件
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	router.POST("/upload", func(c *gin.Context) {
		// 多文件上传
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("F:/BaiduNetdiskDownload/%s_%d", file.Filename, index)
			// 上传文件到指定目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()

}
