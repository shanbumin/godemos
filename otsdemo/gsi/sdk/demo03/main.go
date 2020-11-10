package main

import (
	"otsdemo/gsi/sdk/servers"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)

func main() {
	servers.DeleteIndex(bootstrap.Client, constants.GSI2Table, constants.GSI2Definedcol1Index)
}
