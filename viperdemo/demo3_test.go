package viperdemo

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

// @program:     go-web-exercises
// @file:        demo3_test.go
// @author:      bowen
// @create:      2023-01-09 19:46
// @description: Go语言中文网demo

const (
	//gRpc服务地址
	Address = "0.0.0.0:9090"
)

// 写入配置文件
func TestVipeerW(t *testing.T) {
	viper.SetConfigFile("hello.toml")
	viper.Set("Address", "0.0.0.0:9090")
	if err := viper.WriteConfig(); err != nil {
		panic(fmt.Errorf("Fatal err config file:%s\n", err))
	}
}

// 读取配置文件
func TestRead(t *testing.T) {
	viper.SetConfigFile("hello.toml")
	// 会查找和读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file:%s\n", err))
	}
	Address := viper.GetString("Address")
	//key取Address或者address都能取到值，反正viper转成小写处理
	fmt.Println(Address)
}
