package main

import (
	"otsdemo/bootstrap"
	"otsdemo/sample"
	"otsdemo/table/servers"
)

//过滤器
//在服务端对读取结果再进行一次过滤，根据过滤器（Filter）中的条件决定返回的行。使用过滤器后，只返回符合条件的数据行。

//todo 使用方法
//1.在通过GetRow、BatchGetRow或GetRange接口查询数据时，可以使用过滤器只返回符合条件的数据行。
//2.过滤器目前包括SingleColumnValueFilter和CompositeColumnValueFilter。
//  SingleColumnValueFilter：只判断某个参考列的列值。
//  CompositeColumnValueFilter：根据多个参考列的列值的判断结果进行逻辑组合，决定是否过滤某行。

//todo 注意
//当在该次扫描的5000行或者4 MB数据中没有满足过滤器条件的数据时，得到的Response中的Rows为空，
//但是NextStartPrimaryKey可能不为空，此时需要使用NextStartPrimaryKey继续读取数据，直到NextStartPrimaryKey为空。


func main() {

	servers.GetRowWithFilter(bootstrap.Client,sample.Test4Table)
	servers.GetRowWithCompositeColumnValueFilter(bootstrap.Client,sample.Test4Table)
}
