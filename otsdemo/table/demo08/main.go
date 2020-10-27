package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


//前提条件
//
//已初始化Client，详情请参见初始化。
//已创建数据表。
//todo 已删除数据表上的索引表和多元索引。

func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.删除数据表
	sample.DeleteTableSample(client,"t1")




}
