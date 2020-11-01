package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/globalsecondaryindex/ready/prepare"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)
import "fmt"

func main() {

	//先创建一个初始表
	prepare.CreateGSI1TableSample(start.Client,sample.GSI1)
	//BatchWriteGSI1TableSample(start.Client,sample.GSI1)
	//插入初始化数据
	prepare.BatchWriteGSI1TableSample(start.Client,sample.GSI1)
}













//创建表2
//todo  为了演示创建完数据表之后再创建索引表，所以这里我们创建数据表的时候不创建索引表啦
func CreateTableWithGlobalIndexSample(client *tablestore.TableStoreClient, tableName string) {

	createTableRequest := new(tablestore.CreateTableRequest)
	//1.数据表的结构信息
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddDefinedColumn("definedcol1", tablestore.DefinedColumn_STRING) //添加预定义列1
	tableMeta.AddDefinedColumn("definedcol2", tablestore.DefinedColumn_INTEGER) //添加预定义列2
	createTableRequest.TableMeta = tableMeta




	//1.tableMeta (创建主键列的schema)


	indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的definedcol1列作为索引表的主键。
	indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。
	indexMeta.IndexName = "indexSample"
	//2.tableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1 //数据的过期时间，单位为秒，-1表示永不过期。带索引表的数据表数据生命周期必须设置为-1。
	tableOption.MaxVersion = 1 //保存的最大版本数，1表示每列上最多保存一个版本即保存最新的版本。带索引表的数据表最大版本数必须设置为1。
	//3.reservedThroughput(保留吞吐量)
	reservedThroughput := new(tablestore.ReservedThroughput)

	//4.createtableRequest


	createTableRequest.TableOption = tableOption
	createTableRequest.ReservedThroughput = reservedThroughput
	createTableRequest.AddIndexMeta(indexMeta)
	//5.create
	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}
}


