package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

// 日志分开存储
func TestZap(t *testing.T) {
	//encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 设置时间
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	logFile1, _ := os.Create("app.log")
	core1 := zapcore.NewCore(encoder, zapcore.AddSync(logFile1), zapcore.InfoLevel)

	logFile2, _ := os.Create("app.err.log")
	core2 := zapcore.NewCore(encoder, zapcore.AddSync(logFile2), zapcore.ErrorLevel)

	newCow := zapcore.NewTee(core1, core2)
	logger := zap.New(newCow)

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
