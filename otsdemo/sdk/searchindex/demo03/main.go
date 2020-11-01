package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {
	//查询多元索引描述信息
	sample.DescribeSearchIndex(start.Client,sample.SearchIndexName,sample.SearchIndex1)
}
