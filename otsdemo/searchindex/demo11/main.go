package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//短语匹配查询
	servers.MatchPhraseQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
