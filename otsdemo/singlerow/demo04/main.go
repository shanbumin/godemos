package main

import (
	"otsdemo/bootstrap"
	"otsdemo/sample"
	"otsdemo/singlerow/servers"
)


//DeleteRow接口用于删除一行数据。如果删除的行不存在，则不会发生任何变化。
func main() {
	//删除一行数据
	servers.DeleteRowSample(bootstrap.Client,sample.DemoTable)
}
