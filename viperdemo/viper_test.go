package viperdemo

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestViper(t *testing.T) {
	// 1. 指定配置文件来源，这里除了从路径获取，还有其他很多方法。
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 使用配置
	fmt.Println(viper.Get("port").(int))
	fmt.Println(viper.GetInt("port"))
	fmt.Println(viper.Get("version"))
	fmt.Println(viper.GetString("version"))
	fmt.Println(viper.GetInt("app.node"))
	//viper.Get("app")

	fmt.Println()

}

func TestViperV2(t *testing.T) {
	// 1. 指定配置文件来源，这里除了从路径获取，还有其他很多方法。
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	for t := range time.Tick(time.Second * 3) {
		fmt.Printf("%s:%v\n", t.Format("2006/01/02 15:04:05"), viper.GetString("version"))
	}

}

func TestCallback(t *testing.T) {
	// 1. 指定配置文件来源，这里除了从路径获取，还有其他很多方法。
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变化时会执行回调函数
		fmt.Println("config file changed:", e.Name)
	})
	select {}
}

func callback(e fsnotify.Event) {
	// 配置文件发生变化时会执行回调函数
	//fmt.Println("配置文件发生变化:", e.Name)
	fmt.Printf("配置文件发生变化：%#v\n", e)
	if e.Op == fsnotify.Chmod {
		fmt.Println("修改配置项")
	}
	if e.Op == fsnotify.Create {
		fmt.Println("增加配置项")
	}
	if e.Op == fsnotify.Write {
		fmt.Println("hhh")

	}

}

func TestCallbackV1(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(callback)
	select {}
}

func TestEnv(t *testing.T) {
	viper.SetEnvPrefix("todo")
	viper.BindEnv("name")
	fmt.Println(viper.Get("name"))
}
