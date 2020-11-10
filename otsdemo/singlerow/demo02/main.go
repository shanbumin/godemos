package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/singlerow/servers"
)



//读取的结果可能有如下两种：
//1.如果该行存在，则返回该行的各主键列以及属性列。
//2.如果该行不存在，则返回中不包含行，并且不会报错。
func main() {
	//GetRow接口用于读取一行数据。
	servers.GetRowSample(bootstrap.Client, constants.DemoTable)


}
