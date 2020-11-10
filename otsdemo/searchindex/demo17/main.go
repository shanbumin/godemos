package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/constants"
)

func main() {
	//1.初始化对接
	client:=tablestore.NewClient(constants.EndPoint, constants.InstanceName, constants.AccessKeyId, constants.AccessKeySecret)
	//2.折叠(去重)
	_=client
}
