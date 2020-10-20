package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.查询表描述信息
	//todo 使用DescribeTable接口可以查询指定表的结构、预留读/写吞吐量详情等信息。
	sample.DescribeTableSample(client,"t1")
}