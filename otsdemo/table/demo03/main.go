package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.创建表
	//todo 您可以使用UpdateTable接口来更新指定表的预留读/写吞吐量。
	sample.UpdateTableSample(client,"t1")
}