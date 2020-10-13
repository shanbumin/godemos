package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

/*

bogon:greet sam$ greet  help
NAME:
   greet - fight the loneliness!

USAGE:
   greet [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

 */

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}