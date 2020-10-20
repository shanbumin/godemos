package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient("https://tulong.cn-shanghai.ots.aliyuncs.com",  "tulong", "LTAI4G4nT4uRrmdmpz2XGkpV", "eBRDHMUr1NhuBBlqno1U7Hsi5adZ5O")
	//2.创建表
	//todo 创建数据表（带索引）
	sample.CreateTableWithGlobalIndexSample(client,"t2")
}