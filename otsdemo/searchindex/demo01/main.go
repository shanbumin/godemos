package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {
	//创建多元索引
	sample.CreateSearchIndex(bootstrap.Client,sample.SearchIndexName,sample.SearchIndex1)
}
