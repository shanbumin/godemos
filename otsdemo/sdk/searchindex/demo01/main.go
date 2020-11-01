package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {
	//创建多元索引
	sample.CreateSearchIndex(start.Client,sample.SearchIndexName,sample.SearchIndex1)
}
