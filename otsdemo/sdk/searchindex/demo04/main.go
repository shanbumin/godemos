package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {
	//删除多元索引
	sample.DeleteSearchIndex(start.Client,sample.SearchIndexName,sample.SearchIndex1)
}
