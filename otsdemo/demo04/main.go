package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient("https://tulong.cn-shanghai.ots.aliyuncs.com",  "tulong", "LTAI4G4nT4uRrmdmpz2XGkpV", "eBRDHMUr1NhuBBlqno1U7Hsi5adZ5O")
	//2.列出表名称
	//todo 使用ListTable接口获取当前实例下已创建的所有表的表名。
	sample.ListTableSample(client)
}