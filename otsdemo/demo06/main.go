package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.预定义列操作
	//todo 为数据表增加预定义列或删除数据表的预定义列。
	//todo 设置预定义列后，在创建全局二级索引时将预定义列作为索引表的索引列或者属性列。


	sample.DescribeTableSample(client,"t1")
}