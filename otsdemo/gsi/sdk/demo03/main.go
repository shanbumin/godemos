package main

import (
	"otsdemo/gsi/sdk/servers"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {
	servers.DeleteIndex(start.Client,sample.GSI2Table,sample.GSI2Definedcol1Index)
}
