package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.并发导出数据
	_=client
}
