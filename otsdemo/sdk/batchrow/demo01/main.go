package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)






func main() {
	//批量写
	sample.BatchWriteRowSample(start.Client,sample.BatchName)
}
