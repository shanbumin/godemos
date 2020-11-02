package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)



func main() {
	//批量读
	servers.BatchGetRowSample(bootstrap.Client,sample.BatchName)
	
}
