package main

import "fmt"
import "github.com/robfig/cron"



//Job 类型
//// Job is an interface for submitted cron jobs.
//type Job interface {
//	Run()
//}
//
//由此可知，Job是带有一个Run方法的接口类型，经过代码分析可以指定，cron定时调度时间到达时，将调用此
//方法，也就是意味着，任何实现了Run方法的实例，都可以作为AddJob函数的cmd参数，而Run方法所实现的内容
//就是你定时调度所需执行的任务(todo AddFunc函数只能添加无参数无返回的任务，太鸡肋了)，接下来我们就来实现一个带参数的任务添加



//cron实现带参数的任务添加
//定义一个类型 包含一个int类型参数和函数体
type funcIntJob struct {
	num   int
	function func(int)
}

//实现这个类型的Run()方法 使得可以传入Job接口
func (this *funcIntJob) Run() {
	if nil != this.function {
		this.function(this.num)
	}
}

//非必须  返回一个urlServeJob指针
func newfuncIntJob(num int, function func(int)) *funcIntJob {
	instance := &funcIntJob{
		num:   num,
		function: function,
	}
	return instance
}

//示例任务
func shownum(num int){
	fmt.Println(num)
}

func main(){
	var c = cron.New()
	job := newfuncIntJob(3, shownum)
	spec := "*/5 * * * * ?"
	c.AddJob(spec, job)
	c.Start()
	defer c.Stop()
	select{}
}


