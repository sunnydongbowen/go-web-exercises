package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

// @program:     go-web-exercises
// @file:        zapv7_test.go
// @author:      bowen
// @create:      2022-11-12 11:07
// @description: 日志切割

func TestZapSlpit(t *testing.T) {

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), zapcore.InfoLevel)

	logger := zap.New(core)

	var uid int64 = 189811282
	islogin := true
	name := "bowen"
	logger.Info("日志信息",
		zap.Int64("uid", uid),
		zap.Bool("isLogin", islogin),
		zap.String("name", name),
		zap.Any("data", name))
	// 这条日志不会打印。会有leve check
	logger.Debug("这是一条debug日志")
	logger.Error("这是一条错误日志")

}
