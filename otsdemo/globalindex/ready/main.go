package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)
import "fmt"

func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.先创建一个初始表
	CreateTableWithGlobalIndexSample(client,"sbm_global_index")
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



//func CreateTableWithGlobalIndexSample(client *tablestore.TableStoreClient, tableName string) {
//	fmt.Println("Begin to create table:", tableName)
//	createtableRequest := new(tablestore.CreateTableRequest)
//
//	tableMeta := new(tablestore.TableMeta)
//	tableMeta.TableName = tableName
//	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
//	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
//	tableMeta.AddDefinedColumn("definedcol1", tablestore.DefinedColumn_STRING)
//	tableMeta.AddDefinedColumn("definedcol2", tablestore.DefinedColumn_INTEGER)
//
//	indexMeta := new(tablestore.IndexMeta)
//	indexMeta.AddPrimaryKeyColumn("pk1")
//	indexMeta.AddDefinedColumn("definedcol1")
//	indexMeta.AddDefinedColumn("definedcol2")
//	indexMeta.IndexName = "testindex1"
//
//	tableOption := new(tablestore.TableOption)
//	tableOption.TimeToAlive = -1
//	tableOption.MaxVersion = 1
//	reservedThroughput := new(tablestore.ReservedThroughput)
//	reservedThroughput.Readcap = 0
//	reservedThroughput.Writecap = 0
//	createtableRequest.TableMeta = tableMeta
//	createtableRequest.TableOption = tableOption
//	createtableRequest.ReservedThroughput = reservedThroughput
//
//	createtableRequest.AddIndexMeta(indexMeta)
//
//	_, err := client.CreateTable(createtableRequest)
//
//	if err != nil {
//		fmt.Println("Failed to create table with error:", err)
//	} else {
//		fmt.Println("Create table finished")
//	}
//
//
//	indexMeta.IndexName = "index2"
//	indexReq := &tablestore.CreateIndexRequest{ MainTableName:tableName, IndexMeta: indexMeta, IncludeBaseData: false }
//	resp, err := client.CreateIndex(indexReq)
//	if err != nil {
//		fmt.Println("Failed to create table with error:", err)
//	} else {
//		fmt.Println("Create index finished", resp)
//	}
//
//	deleteIndex := &tablestore.DeleteIndexRequest{ MainTableName:tableName, IndexName: indexMeta.IndexName }
//	resp2, err := client.DeleteIndex(deleteIndex)
//
//	if err != nil {
//		fmt.Println("Failed to create table with error:", err)
//	} else {
//		fmt.Println("drop index finished", resp2)
//	}
//
//	describeTableReq := new(tablestore.DescribeTableRequest)
//	describeTableReq.TableName = tableName
//	describ, err := client.DescribeTable(describeTableReq)
//
//	if err != nil {
//		fmt.Println("failed to update table with error:", err)
//	} else {
//		fmt.Println("DescribeTableSample. indexinfo:", describ.IndexMetas[0], len(describ.IndexMetas))
//	}
//
//	addColumnsReq := new(tablestore.AddDefinedColumnRequest)
//	addColumnsReq.TableName = tableName
//	addColumnsReq.AddDefinedColumn("definedcol3", tablestore.DefinedColumn_INTEGER)
//
//	_, err = client.AddDefinedColumn(addColumnsReq)
//	if err != nil {
//		fmt.Println("failed to add defined column with error:", err)
//	}
//
//	describeTableReq.TableName = tableName
//	describ, err = client.DescribeTable(describeTableReq)
//
//	if err != nil {
//		fmt.Println("failed to describe table with error:", err)
//	} else {
//		fmt.Println(describ.TableMeta.DefinedColumns[0].Name, describ.TableMeta.DefinedColumns[0].ColumnType)
//		fmt.Println(describ.TableMeta.DefinedColumns[1].Name, describ.TableMeta.DefinedColumns[1].ColumnType)
//		fmt.Println(describ.TableMeta.DefinedColumns[2].Name, describ.TableMeta.DefinedColumns[2].ColumnType)
//		fmt.Println("DescribeTableSample finished. indexinfo:", describ.IndexMetas[0], len(describ.IndexMetas))
//	}
//}
