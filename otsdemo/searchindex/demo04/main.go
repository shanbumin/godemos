package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//删除多元索引
	servers.DeleteSearchIndex(bootstrap.Client, constants.DemoTable, constants.DemoTableIndex)
}
