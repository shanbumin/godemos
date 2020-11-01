package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)


//创建测试表sbm_demos
//todo 创建数据表的时候是可以同时创建全局二级索引表的，但多元索引表需要后期创建
//todo  全局二级索引是创建数据表的一个组成项(属性)
func main() {
	    tableName:=sample.DemoTableName
	    //gsindex01:="gs_name_index"
	    client:= start.Client
		createTableRequest := new(tablestore.CreateTableRequest)
		//1.TableMeta
		tableMeta := new(tablestore.TableMeta)
		tableMeta.TableName = tableName
		tableMeta.AddPrimaryKeyColumn("_id", tablestore.PrimaryKeyType_STRING)
	    tableMeta.AddDefinedColumn("name", tablestore.DefinedColumn_STRING) //添加预定义列  姓名
	    tableMeta.AddDefinedColumn("age", tablestore.DefinedColumn_INTEGER) //年龄
	    tableMeta.AddDefinedColumn("salary",tablestore.DefinedColumn_DOUBLE) //薪水
	    tableMeta.AddDefinedColumn("married",tablestore.DefinedColumn_BOOLEAN) //是否已婚
	    tableMeta.AddDefinedColumn("desc",tablestore.DefinedColumn_BINARY) //描述 todo 可以用来存图片二进制流的额，这里我们仅仅是演示而已
	    tableMeta.AddDefinedColumn("created_at",tablestore.DefinedColumn_INTEGER) //创建时间
	    tableMeta.AddDefinedColumn("updated_at",tablestore.DefinedColumn_INTEGER) //修改时间
		createTableRequest.TableMeta = tableMeta

		//2.IndexMeta
		//indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
		//indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的definedcol1列作为索引表的主键。
		//indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的definedcol1列作为索引表的主键。
		//indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。
		//indexMeta.IndexName = gsindex01
		//createTableRequest.AddIndexMeta(indexMeta)

	    //3.TableOption
		tableOption := new(tablestore.TableOption)
		tableOption.TimeToAlive = -1
		tableOption.MaxVersion = 1
		createTableRequest.TableOption = tableOption

		//4.ReservedThroughput
		reservedThroughput := new(tablestore.ReservedThroughput)
		reservedThroughput.Readcap = 0
		reservedThroughput.Writecap = 0
		createTableRequest.ReservedThroughput = reservedThroughput





		
		//CreateTable
		_, err := client.CreateTable(createTableRequest)
		if err != nil {
			fmt.Println("Failed to create table with error:", err)
		} else {
			fmt.Println("Create table finished")
		}
}
