package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//列存在性查询
	servers.ExistQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
