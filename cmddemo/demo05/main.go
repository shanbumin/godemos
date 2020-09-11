package main

import (
	"fmt"
	"os"
	"time"
)

//golang 进程属性和控制
//@link https://blog.csdn.net/qq_21816375/article/details/78315303
func main() {

	//1.进程ID
	pid:=os.Getpid()
	fmt.Println("当前进程ID:",pid)

	ppid:=os.Getppid()
	fmt.Println("当前父进程ID:",ppid)

	//2.可通过 os.Getuid() 和 os.Getgid() 获取当前进程的实际用户 ID 和实际组 ID；
	fmt.Println("-----------------------------")
    uid:=os.Getuid()
	gid:=os.Getgid()
	fmt.Println("当前进程实际用户ID:",uid)
	fmt.Println("当前进程实际组ID:",gid)

	time.Sleep(60 * time.Second)

}
