package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {

	//全匹配查询
	servers.MatchAllQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
