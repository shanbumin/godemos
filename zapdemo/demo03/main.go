package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {

	//1.EncoderConfig,比如json key 自定义、时间格式化
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	//2.设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	//3.构建日志logger
	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      false,                                                // 开发模式，堆栈跟踪
		Encoding:         "console",                                              // 输出格式 console 或 json, todo console一种格式不是说要输出到控制台
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
		//OutputPaths:      []string{"stdout", "./logs/spikeProxy.log"},         // 这里配置到标准输出以及文件中
		OutputPaths:      []string{"stdout"},     //配置多个，就是输出到不同的地方了，这里仅配置输出到标准输出
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}


	//4.调试
	logger.Info("log 初始化成功")
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Error("fuck",zap.String("say","go away"))
}