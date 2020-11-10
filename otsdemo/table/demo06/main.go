package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/table/servers"
)


//前提条件
//todo 已删除数据表上的索引表和多元索引。

func main() {

	//删除数据表
	servers.DeleteTableSample(bootstrap.Client, constants.TestTable)




}
