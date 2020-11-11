package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {

	//嵌套类型查询
	servers.NestedQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
