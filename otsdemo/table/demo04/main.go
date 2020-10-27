package main

import (
	"otsdemo/sample"
	"otsdemo/start"
)


func main() {

	//查询表描述信息
	//todo 使用DescribeTable接口可以查询指定表的结构、预留读/写吞吐量详情等信息。
	sample.DescribeTableSample(start.Client,sample.TableName)
}