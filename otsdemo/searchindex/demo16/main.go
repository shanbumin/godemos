package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

//使用多元索引时，可以在创建时指定索引预排序和在查询时指定排序方式，在获取返回结果时可以使用Limit和Offset或者使用Token进行翻页。
func main() {
	//排序和翻页
	servers.QueryRowsWithToken(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
