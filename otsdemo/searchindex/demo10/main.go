package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)
//RangeQuery根据范围条件查询表中的数据。对于Text类型字段，只要分词后的词条中有词条满足范围条件即可。
func main() {
	//范围查询
	servers.RangeQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
