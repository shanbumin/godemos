package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

//PrefixQuery根据前缀条件查询表中的数据。对于Text类型字段，只要分词后的词条中有词条满足前缀条件即可。

func main() {
	  servers.PrefixQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
