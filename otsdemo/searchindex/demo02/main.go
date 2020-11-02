package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {

	//列出多元索引列表
	sample.ListSearchIndex(bootstrap.Client,sample.SearchIndexName)
}
