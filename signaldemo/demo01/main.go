package main

import (
	"fmt"
	"os"
	"os/signal"
)


//todo 在调试的时候，当执行go run main.go的时候，千万别通过ps aux|grep main.go来获取当前进程的id
//todo go run main.go本质是会发起一个子进程自动编译成二进制文件执行的，所以执行的程序应该是那个二进制才对
//todo 可以通过这种方式查看: ps aux|grep "go-build"  或者 ps aux|grep "/var/folders"
//todo 或者我们直接在程序中打印出也可以的额  fmt.Println(os.Getpid())

// 监听全部信号
func main()  {

	fmt.Println(os.Getpid())
	 
	//创建chan
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)
	//阻塞直到有信号传入
	fmt.Println("启动")
	s := <-c
	fmt.Println("退出信号", s)
}

//go run main.go
//1.ctrl+c退出,输出
// kill pid
// kill -30 pid
// kill -31 pid
// kill -9  pid #执行默认操作