package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)






func main() {
	//批量写
	servers.BatchWriteRowSample(bootstrap.Client, constants.BatchTable)
}
