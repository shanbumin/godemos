package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)




//pk1   string
//pk2   integer
//pk3   binary


func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.插入一行数据（PutRow）
	sample.PutRowSample(client,"t12")
}
