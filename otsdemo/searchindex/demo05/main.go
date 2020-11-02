package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)


//todo 前提条件

//已创建数据表并写入数据。
//已在数据表上创建多元索引。
func main() {
	//精确查询
	sample.TermQuery(bootstrap.Client,sample.SearchIndexName,sample.SearchIndex1)
}
