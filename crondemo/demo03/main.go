package main

import (
	"github.com/robfig/cron"
	"log"
)


type TestJob struct {
}

func (this TestJob)Run() {
	log.Println("testJob1...")
}

//-------------------------------------------------

type Test2Job struct {
}

func (this Test2Job)Run() {
	log.Println("testJob2...")
}

//启动多个任务
func main() {
	i := 0
	c := cron.New()

	//AddFunc
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	//AddJob方法  一个定时任务绑定了多个任务
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select{}
}