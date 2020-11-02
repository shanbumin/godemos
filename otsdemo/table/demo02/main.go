package main

import (
	"otsdemo/bootstrap"
	"otsdemo/sample"
	"otsdemo/table/servers"
)


func main() {
	//更新表
	//todo 您可以使用UpdateTable接口来更新指定表的预留读/写吞吐量。
	servers.UpdateTableSample(bootstrap.Client,sample.TestTable)
}