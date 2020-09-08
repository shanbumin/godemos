package main

import (
	"os/exec"
	"context"
	"time"
	"fmt"
)

type result struct {
	err error
	output []byte
}

func main() {
	//  执行1个cmd, 让它在一个协程里去执行, 让它执行2秒: sleep 2; echo sam;  1秒的时候, 我们杀死cmd
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	// 创建了一个结果队列
	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())
    //异步开启
	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo sam;")

		// 执行任务, 捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果, 传给主协程
		resultChan <- &result{
			err: err,
			output: output,
		}
	}()

	//故意睡1s
	time.Sleep(1 * time.Second)
	// 取消上下文,强杀该命令行的执行
	cancelFunc()
	// 在main协程里, 等待子协程的退出，并打印任务执行结果
	res = <- resultChan
	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
