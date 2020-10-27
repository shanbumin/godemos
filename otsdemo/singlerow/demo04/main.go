package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


//DeleteRow接口用于删除一行数据。如果删除的行不存在，则不会发生任何变化。
func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.删除一行数据
	sample.DeleteRowSample(client,"t12")
}
