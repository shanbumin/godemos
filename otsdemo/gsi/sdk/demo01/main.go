package main

import (
	"otsdemo/gsi/sdk/servers"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {

	//创建索引表
	//使用CreateIndex接口在已存在的数据表上创建索引表。
	//todo 说明 您也可以使用CreateTable接口在创建数据表的同时创建一个或者多个索引表。
	servers.CreateGlobalIndexSample(start.Client,sample.GSI2Table)
}
