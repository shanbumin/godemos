package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


//UpdateRow接口用于更新一行数据，可以增加和删除一行中的属性列，删除属性列指定版本的数据，或者更新已存在的属性列的值。
//如果更新的行不存在，则新增一行数据。
//todo 当UpdateRow请求中只包含删除指定的列且该行不存在时，则该请求不会新增一行数据。
func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.更新一行数据（UpdateRow）
	sample.UpdateRowSample(client,"t12")


}
