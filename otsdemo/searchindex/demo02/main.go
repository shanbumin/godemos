package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {

	//列出多元索引列表
	servers.ListSearchIndex(bootstrap.Client, constants.DemoTable)
}
