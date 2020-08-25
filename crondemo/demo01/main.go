package main
import (
	"github.com/robfig/cron"
	"time"
)
import "fmt"


//cron 表达式代表了一个时间集合，使用 6 个空格分隔的字段表示。

//字段名	是否必须	允许的值	允许的特定字符
//秒(Seconds)	是	0-59	* / , -
//分(Minutes)	是	0-59	* / , -
//时(Hours)	是	0-23	* / , -
//日(Day of month)	是	1-31	* / , – ?
//月(Month)	是	1-12 or JAN-DEC	* / , -
//星期(Day of week)	否	0-6 or SUM-SAT	* / , – ?


func main() {

	c := cron.New()
	//注册在指定时间上运行的函数，cron 将会在协程中运行这些注册函数。AddFunc 函数第一个参数指定定时任务时间间隔，第二个参数指定运行函数。
	c.AddFunc("*/1 * * * *", func() { fmt.Println("Every second.") })
	c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	//上述添加每个Func都会异步的在他们自己的G上调用额
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	//检查cron作业条目的下一个和上一个运行时间。
	//inspect(c.Entries())
	//停止定时任务(不停止已经运行的任务)
	time.Sleep(120 * time.Second)
	c.Stop()

	//Parse()：解析与校验cron表达式
	//Next():根据当前时间，计算下次调度时间




	//调度多个cron任务的方案：
	//
	//1.首先有个map保存了cron任务列表   cron1:10:05    cron2:09:59     cron3:10:00
	//2.存在一个调度协程:
	//a.根据当前时间10:00 去检查过期或者到期的任务(执行cron1协程，执行cron3协程)
	//b.更新对应的下次时间
	//c.休息一会，然后继续调度工作


}
