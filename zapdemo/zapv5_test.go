package main

import (
	"go.uber.org/zap"
	"testing"
)

// @program:     go-web-exercises
// @file:        zapv5_test.go
// @author:      bowen
// @create:      2022-11-10 21:38
// @description: 51job zap练习

func TestZapSugar(t *testing.T) {
	sugar := zap.NewExample().Sugar()
	sugar.Infof("hello! name:%s,age:%d", "xiaomin", 20)

	logger := zap.NewExample()
	logger.Info("hello",
		zap.String("name", "小明"),
		zap.Int64("age", 20))
}

func TestZapCreate(t *testing.T) {
	logger := zap.NewExample()
	logger.Info("example")

	logger, _ = zap.NewDevelopment()
	logger.Info("development")
	//支持结构化
	logger, _ = zap.NewProduction()
	logger.Info("Production")
}
