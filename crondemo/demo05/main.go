package main

import (
	"fmt"
	"time"
)
import "github.com/robfig/cron"

func main() {


	sch, err := cron.Parse("*/5 * * * * *")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sch.Next(time.Now()))



	
}
