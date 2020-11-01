package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {
	//先创建一个测试表
	sample.CreateDemoTable(start.Client,sample.SearchIndexName)

}
