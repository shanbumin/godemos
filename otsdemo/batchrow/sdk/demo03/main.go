package main

import (
	"otsdemo/batchrow/prepare"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)




func main() {
	//范围读
	prepare.GetRangeSample(start.Client,sample.BatchName)
}
