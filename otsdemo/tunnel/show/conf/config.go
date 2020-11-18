package conf

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

//通道回退配置
var DefaultBackoffConfig = tunnel.ChannelBackoffConfig{
	MaxDelay:  5 * time.Second,
	//baseDelay: 20 * time.Millisecond,
	//factor:    5,
	//jitter:    0.25,
}


//日志切割
var DefaultSyncer = zapcore.AddSync(&lumberjack.Logger{
	Filename:   "tunnelClient.log",
	MaxSize:    512, //MB
	MaxBackups: 5,
	MaxAge:     30, //days
	Compress:   true,
})

//日志整体配置
var DefaultLogConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),//日志级别
	Development: false,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	},
	Encoding: "console",
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
}



