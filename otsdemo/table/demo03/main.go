package main

import (
	"otsdemo/bootstrap"
	"otsdemo/table/servers"
)


func main() {

	//列出表名称
	//todo 使用ListTable接口获取当前实例下已创建的所有表的表名。
	servers.ListTableSample(bootstrap.Client)
}