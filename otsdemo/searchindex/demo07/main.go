package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

//类似于TermQuery，但是TermsQuery可以指定多个查询关键词，查询匹配这些词的数据。多个查询关键词中只要有一个词精确匹配，该行数据就会被返回，等价于SQL中的In。
func main() {
	//多词精确查询
	servers.TermsQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
