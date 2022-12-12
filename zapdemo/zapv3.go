package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	//encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	file, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWs := zapcore.AddSync(file)
	consoleWs := zapcore.AddSync(os.Stdout)
	//　这里要转一下类型，不然传给newcore的时候会报错。
	level, err := zapcore.ParseLevel(GetLevel())

	if err != nil {
		level = zapcore.InfoLevel
	}
	//fmt.Println(level)

	// 只会打印info级别以及以上的！
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWs, consoleWs), level)

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
