package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

//加入子命令
func main() {
	log.Printf("Main start!\n")
	//1.创建命令行应用
	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "make an explosive entrance"

	//2.设置子命令
	runCommand := cli.Command{
		Name: "run",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "it",
				Usage: "enable tty",
			},
		},
		Action: func(c *cli.Context) error {
			log.Printf("runcommand args:%s\n", c.Args())
			log.Printf("runcommand tty:%v\n", c.Bool("it"))
			for i := 0; i < 5; i++ {
				log.Printf("runcommand sleep %d\n", i)
				time.Sleep(1 * time.Second)
			}
			return nil
		},
	}

	app.Commands = []cli.Command {
		runCommand,
	}

	//3.行动
	app.Action = func(c *cli.Context) error {
		log.Printf("main function args:%s\n", c.Args())
		for i := 0; i < 5; i++ {
			log.Printf("main function sleep %d\n", i)
			time.Sleep(1 * time.Second)
		}
		return nil
	}

	//4.执行
	log.Printf("Before invoking Run!\n")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main end!\n")
}