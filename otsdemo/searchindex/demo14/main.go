package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)
//BoolQuery查询条件包含一个或者多个子查询条件，根据子查询条件来判断一行数据是否满足查询条件。每个子查询条件可以是任意一种Query类型，包括BoolQuery。
func main() {
	//多条件组合查询
	servers.BoolQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
