package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {
	//先创建一个测试表
	sample.CreateDemoTable(bootstrap.Client,sample.SearchIndexName)

}
