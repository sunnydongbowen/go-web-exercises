package main

import "github.com/spf13/viper"

func GetLevel() string {
	// 1. 指定配置文件来源，这里除了从路径获取，还有其他很多方法。
	viper.SetConfigFile("zapdemo/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	level := viper.GetString("log.level")
	return level
}
