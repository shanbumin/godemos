package main

import (
	"otsdemo/sample"
	"otsdemo/bootstrap"
)




//pk1   string
//pk2   integer
//pk3   binary


func main() {
	//插入一行数据（PutRow）
	sample.PutRowSample(bootstrap.Client,sample.TableConditionName)
}
