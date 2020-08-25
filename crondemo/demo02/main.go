package main

import "log"
import "github.com/robfig/cron"
func main() {


	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i+=5
		log.Println("cron running:", i)
	})
	c.Start()

	select{}



}
