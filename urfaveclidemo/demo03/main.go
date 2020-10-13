package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

//todo  命令行参数
//go run main.go  crontab
func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}