package main

import (
	"go.uber.org/zap"
	"testing"
)

func TestZapv1(t *testing.T) {
	// 获取logger对象
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	var uid int64 = 189811282
	islogin := true
	name := "bowen"
	logger.Info("日志信息",
		zap.Int64("uid", uid),
		zap.Bool("isLogin", islogin),
		zap.String("name", name),
		zap.Any("data", name))

	slogger := logger.Sugar()
	slogger.Info("sugar logger 记录日志 ", uid, islogin, name)
	slogger.Errorf("sugar logger记录日志 ,uid:%v islogin:%v name:%v", uid, islogin, name)
}
