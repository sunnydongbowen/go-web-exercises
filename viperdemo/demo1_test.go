package viperdemo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"testing"
)

func TestDemo1(t *testing.T) {
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()        // 读取配置信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监控配置信息文件变化
	viper.WatchConfig()

	r := gin.Default()

	// 访问/version 的返回值会随着配置文件的变化而变化

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	portInt := viper.Get("port").(int)
	port := ":" + strconv.Itoa(portInt)
	r.Run(port)

	//if err:=r.Run(fmt.Sprintf("%d")),
	//if err := r.Run(
	//	fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
	//	panic(err)
	//}

}
