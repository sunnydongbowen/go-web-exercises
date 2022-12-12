package logger

// @program:     go-web-exercises
// @file:        logger.go
// @author:      bowen
// @create:      2022-11-12 11:35
// @description:
import (
	"go-web-exercises/gin_zap/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var Logger *zap.Logger

// Init 初始化日志
func Init() error {
	// 1. encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	// encoder := zapcore.NewConsoleEncoder(encoderCfg)
	// 2. writesyncer
	lumberJackLogger := &lumberjack.Logger{
		Filename:   setting.Conf.LogConfig.Filename,
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge,
		Compress:   false,
	}
	//这是为了error日志，新弄了一个err日志文件
	lumberJackLogger2 := &lumberjack.Logger{
		Filename:   strings.Replace(setting.Conf.LogConfig.Filename, ".log", "_err.log", 1),
		MaxSize:    setting.Conf.LogConfig.MaxSize,
		MaxBackups: setting.Conf.LogConfig.MaxBackups,
		MaxAge:     setting.Conf.LogConfig.MaxAge,
		Compress:   false,
	}
	level, err := zapcore.ParseLevel(setting.Conf.LogConfig.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}
	// 创建zapcore
	core1 := zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger), level)

	// 这里只存放err日志
	core2 := zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger2), zapcore.ErrorLevel)

	core := zapcore.NewTee(core1, core2)
	// 利用core生成logger
	lg := zap.New(core, zap.AddCaller())
	// Logger = lg
	// 替换zap全局的logger
	zap.ReplaceGlobals(lg)
	zap.L().Info("logger init success")
	return nil
}
