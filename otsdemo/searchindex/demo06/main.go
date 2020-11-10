package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)

func main() {
	//1.初始化对接
	client:=tablestore.NewClient(constants.EndPoint, constants.InstanceName, constants.AccessKeyId, constants.AccessKeySecret)
	//2.全匹配查询
	servers.MatchAllQuery(client,"sbm_search_index","idx001")
}
