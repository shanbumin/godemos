package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)



func main() {
	//批量读
	servers.BatchGetRowSample(bootstrap.Client, constants.BatchTable)
	
}
