package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)



//读取的结果可能有如下两种：
//如果该行存在，则返回该行的各主键列以及属性列。
//如果该行不存在，则返回中不包含行，并且不会报错。
func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.GetRow接口用于读取一行数据。
	sample.GetRowSample(client,"t12")


}
