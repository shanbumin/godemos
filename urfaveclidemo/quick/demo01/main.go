package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

//命令行参数
func main() {
	log.Printf("Main start!\n")
	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		log.Printf("args:%s\n", c.Args())
		for i := 0; i < 5; i++ {
			log.Printf("sleep %d\n", i)
			time.Sleep(1 * time.Second)
		}
		return nil
	}

	log.Printf("Before invoking Run!\n") //执行Run函数前
	//todo 可以看到app.Run(os.Args)是会阻塞的.直到Action执行完毕
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main end!\n")
}
/*
bogon:demo01 sam$ go run main.go  args01 args02


2020/10/13 14:58:49 Main start!
2020/10/13 14:58:49 Before invoking Run!
2020/10/13 14:58:49 args:[args01 args02]
2020/10/13 14:58:49 sleep 0
2020/10/13 14:58:50 sleep 1
2020/10/13 14:58:51 sleep 2
2020/10/13 14:58:52 sleep 3
2020/10/13 14:58:53 sleep 4
2020/10/13 14:58:54 Main end!



 */



