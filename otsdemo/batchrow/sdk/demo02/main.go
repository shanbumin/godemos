package main

import (
	"otsdemo/batchrow/prepare"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)



func main() {
	//批量读
	prepare.BatchGetRowSample(start.Client,sample.BatchName)
	
}
