package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)


//DeleteRow接口用于删除一行数据。如果删除的行不存在，则不会发生任何变化。
func main() {
	//删除一行数据
	sample.DeleteRowSample(start.Client,sample.TableConditionName)
}
