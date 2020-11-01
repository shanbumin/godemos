package main

import (
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

//todo 这里演示创建数据表之后再添加预定义列
func main() {
	//预定义列操作
	//todo 为数据表增加预定义列或删除数据表的预定义列。
	//todo 设置预定义列后，在创建全局二级索引时将预定义列作为索引表的索引列或者属性列。
	sample.AddDefinedColumn(start.Client,sample.TableName)
	//删除预定义列
	sample.DeleteDefinedColumn(start.Client,sample.TableName)


}