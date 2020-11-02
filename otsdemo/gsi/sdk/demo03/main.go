package main

import (
	"otsdemo/gsi/sdk/servers"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

func main() {
	servers.DeleteIndex(bootstrap.Client,sample.GSI2Table,sample.GSI2Definedcol1Index)
}
