package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {
	//删除多元索引
	sample.DeleteSearchIndex(bootstrap.Client,sample.SearchIndexName,sample.SearchIndex1)
}
