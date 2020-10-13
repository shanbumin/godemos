package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

//加入Flag 命令行选项


func main() {
	log.Printf("Main start!\n")
	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "make an explosive entrance"

	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "flag",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name: "lang",
			Value: "english",
		},
	}

	app.Action = func(c *cli.Context) error {
		log.Printf("args:%s\n", c.Args())
		log.Printf("flag:%v\n", c.Bool("flag"))
		log.Printf("lang:%s\n", c.String("lang"))
		for i := 0; i < 5; i++ {
			log.Printf("sleep %d\n", i)
			time.Sleep(1 * time.Second)
		}
		return nil
	}

	log.Printf("Before invoking Run!\n")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main end!\n")
}

/*
bogon:demo02 sam$ go run main.go  --flag  --lang chinese   args01   args02
2020/10/13 15:11:21 Main start!
2020/10/13 15:11:21 Before invoking Run!
2020/10/13 15:11:21 args:[args01 args02]
2020/10/13 15:11:21 flag:true
2020/10/13 15:11:21 lang:chinese
2020/10/13 15:11:21 sleep 0
2020/10/13 15:11:22 sleep 1
2020/10/13 15:11:23 sleep 2
2020/10/13 15:11:24 sleep 3
2020/10/13 15:11:25 sleep 4
2020/10/13 15:11:26 Main end!

 */

