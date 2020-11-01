package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

func main() {


	createTableRequest := new(tablestore.CreateTableRequest)
	//1.TableMeta
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = sample.BatchName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	createTableRequest.TableMeta = tableMeta

	//2.TableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 1
	createTableRequest.TableOption = tableOption

	//3.ReservedThroughput
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createTableRequest.ReservedThroughput = reservedThroughput



	_,err:=start.Client.CreateTable(createTableRequest)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("创建表成功")
	}





}
