package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.创建索引表
	//使用CreateIndex接口在已存在的数据表上创建索引表。
	//todo 说明 您也可以使用CreateTable接口在创建数据表的同时创建一个或者多个索引表。
	sample.CreateGlobalIndexSample(client,"sbm_global_index")
}
