package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/singlerow/servers"
)



func main() {
	//插入一行数据（PutRow）
	servers.PutRowSample(bootstrap.Client, constants.DemoTable)
}
