package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)


func main() {
	//更新表
	//todo 您可以使用UpdateTable接口来更新指定表的预留读/写吞吐量。
	sample.UpdateTableSample(start.Client,sample.TableName)
}