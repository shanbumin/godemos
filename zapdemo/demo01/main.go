package main

import (
	"go.uber.org/zap"
	"time"
)

//todo  2020-11-18T10:59:58.556+0800    INFO    demo01/main.go:12       来客人了:       {"name": "sam", "age": 18, "time": "1s"}
func main() {
	// zap.NewDevelopment 格式化输出
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("来客人了:",
		zap.String("name", "sam"),
		zap.Int("age",18),
		zap.Duration("time", time.Second),
	)
}