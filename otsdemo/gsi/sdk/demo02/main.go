package main

import (
	"otsdemo/gsi/sdk/servers"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

//从索引表中单行或者范围读取数据，当返回的属性列在索引表中时，可以直接读取索引表获取数据，否则请自行反查数据表获取数据。
func main() {
	servers.GetRangeFromIndex(start.Client,sample.GSI2Definedcol1Index)
}
