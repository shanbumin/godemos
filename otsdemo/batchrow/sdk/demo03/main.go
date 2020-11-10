package main

import (
	"otsdemo/batchrow/servers"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)




func main() {
	//范围读
	servers.GetRangeSample(bootstrap.Client, constants.BatchTable)
}
