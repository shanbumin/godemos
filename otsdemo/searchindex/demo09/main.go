package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

//MatchQuery采用近似匹配的方式查询表中的数据。对Text类型的列值和查询关键词会先按照设置好的分词器做切分，然后按照切分好后的词去查询。
//todo 例如某一行数据的title列的值是“杭州西湖风景区”，使用单字分词，如果MatchQuery中的查询词是“湖风”，则可以匹配到该行数据。
//场景
//匹配查询一般应用于全文检索场景，可应用于Text类型。
func main() {

	//匹配查询
	servers.MatchQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
