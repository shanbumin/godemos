package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//折叠(去重)
	servers.QueryWithCollapse(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
