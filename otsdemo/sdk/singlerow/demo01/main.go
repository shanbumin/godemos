package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)




//pk1   string
//pk2   integer
//pk3   binary


func main() {
	//插入一行数据（PutRow）
	sample.PutRowSample(start.Client,sample.TableConditionName)
}
