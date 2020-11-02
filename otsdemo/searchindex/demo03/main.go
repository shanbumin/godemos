package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {
	//查询多元索引描述信息
	sample.DescribeSearchIndex(bootstrap.Client,sample.SearchIndexName,sample.SearchIndex1)
}
