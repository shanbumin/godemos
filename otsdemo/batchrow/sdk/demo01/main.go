package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)






func main() {
	//批量写
	servers.BatchWriteRowSample(bootstrap.Client,sample.BatchName)
}
