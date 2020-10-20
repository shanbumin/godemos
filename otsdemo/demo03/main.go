package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient("https://tulong.cn-shanghai.ots.aliyuncs.com",  "tulong", "LTAI4G4nT4uRrmdmpz2XGkpV", "eBRDHMUr1NhuBBlqno1U7Hsi5adZ5O")
	//2.创建表
	//todo 您可以使用UpdateTable接口来更新指定表的预留读/写吞吐量。
	sample.UpdateTableSample(client,"t1")
}