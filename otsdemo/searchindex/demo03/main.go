package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//查询多元索引描述信息
	servers.DescribeSearchIndex(bootstrap.Client, constants.DemoTable, constants.DemoTableIndex)
}
