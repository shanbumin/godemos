package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)




func main() {
	//范围读
	servers.GetRangeSample(bootstrap.Client,sample.BatchName)
}
