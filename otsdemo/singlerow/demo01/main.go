package main

import (
	"otsdemo/bootstrap"
	"otsdemo/sample"
	"otsdemo/singlerow/servers"
)



func main() {
	//插入一行数据（PutRow）
	servers.PutRowSample(bootstrap.Client,sample.DemoTable)
}
