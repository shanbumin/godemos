package main

import (
	"fmt"
	"strconv"
	"time"
)

import "github.com/robfig/cron"

func main() {
	//返回某个任务的下次调度时间，如果为0，则任务将不会运行
	//注意一个任务可能绑定了多个定时器规则，我们需要循环计算出最近的那次额

	   rules:=[]string{"*/30 * * * * *","*/25 * * * * *","*/5 * * * * *"}
	   fmt.Println("当前时间：",time.Now().Format("2006-01-02 15:04:05"))
	    nextTime := time.Time{}
		for i, r := range rules {
			//解析规则
			sch,_:= cron.Parse(r)
			//将当前时间传递给到Next方法就可以获得该规则的下次调度时间
			t := sch.Next(time.Now())
			fmt.Println("规则"+strconv.Itoa(i)+"下次调度时间：",t.Format("2006-01-02 15:04:05"))
			if i == 0 || t.UnixNano() < nextTime.UnixNano() {
			        nextTime = t
		    }
	    }
	   fmt.Println("该任务的最终下次调度时间：",nextTime.Format("2006-01-02 15:04:05"))
}
