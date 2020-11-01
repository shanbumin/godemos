package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {

	//列出多元索引列表
	sample.ListSearchIndex(start.Client,sample.SearchIndexName)
}
