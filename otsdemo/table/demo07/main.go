package main

import (
	"otsdemo/bootstrap"
	"otsdemo/sample"
	"otsdemo/table/servers"
)



func main() {

	//主键列自增
	//todo 设置非分区键(第一个主键为分区键)的主键列为自增列后，在写入数据时，无需为自增列设置具体值，表格存储会自动生成自增列的值。该值在分区键级别唯一且严格递增。

	//todo 使用方法
	//1.创建表时，将非分区键的主键列设置为自增列。
	//  只有整型的主键列才能设置为自增列，系统自动生成的自增列值为64位的有符号整型。
	//2.写入数据时，无需为自增列设置具体值，只需将相应主键指定为自增主键。
	//  如果需要获取写入数据后系统自动生成的自增列的值，将ReturnType设置为RT_PK，可以在数据写入成功后返回自增列的值。
	//3.查询数据时，需要完整的主键值。通过设置PutRow、UpdateRow或者BatchWriteRow中的ReturnType为RT_PK可以获取完整的主键值。

	servers.CreateTableKeyAutoIncrementSample(bootstrap.Client,sample.Test3Table) //创建表
	servers.PutRowWithKeyAutoIncrementSample(bootstrap.Client,sample.Test3Table)  //写数据

}
