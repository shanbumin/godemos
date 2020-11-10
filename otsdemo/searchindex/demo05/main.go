package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)


//todo 前提条件

//已创建数据表并写入数据。
//已在数据表上创建多元索引。
func main() {
	//精确查询
	servers.TermQuery(bootstrap.Client, constants.DemoTable, constants.DemoTableIndex)
}
