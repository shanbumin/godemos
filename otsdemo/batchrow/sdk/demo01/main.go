package main

import (
	"otsdemo/batchrow/prepare"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)






func main() {
	//批量写
	prepare.BatchWriteRowSample(start.Client,sample.BatchName)
}
