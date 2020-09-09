package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)




//todo golang中多goroutine时kill信号接收的问题(一)
//@link  https://blog.csdn.net/zhuxinquan61/article/details/89421634


// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	fmt.Println("p-",os.Getpid())
	for i := 0; ; i++ {
		out <- i*factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	fmt.Println("c-",os.Getpid())
	for v := range in {
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}
}



//我们可以让`main`函数保存阻塞状态不退出，只有当用户输入`Ctrl-C`时才真正退出程序：

func main() {
   ch := make(chan int, 64) // 成果队列
   go Producer(3, ch) // 生成 3 的倍数的序列
   go Producer(5, ch) // 生成 5 的倍数的序列
   go Consumer(ch)    // 消费 生成的队列

	fmt.Println(os.Getpid())
   // Ctrl+C 退出
   sig := make(chan os.Signal,1)
   signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM,syscall.SIGUSR1, syscall.SIGUSR2)
   fmt.Printf("退出: (%v)\n", <-sig)
   os.Exit(400)
}

