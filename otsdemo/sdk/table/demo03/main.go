package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)


func main() {

	//列出表名称
	//todo 使用ListTable接口获取当前实例下已创建的所有表的表名。
	sample.ListTableSample(start.Client)
}