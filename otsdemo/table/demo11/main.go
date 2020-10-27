package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"log"
	"otsdemo/sample"
)


//todo 根据指定主键删除一行，如果指定的行存在，且col2列的值为3，则更新成功，否则更新失败。
func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.根据指定主键更新一行，如果指定的行存在，则更新成功，否则更新失败。
	tableName:="t11"


	//a.PrimaryKey
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	deletePk.AddPrimaryKeyColumn("pk2", int64(2))
	deletePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))


    //b.DeleteRowRequest
	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = tableName
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)//期望行存在。

	//CT_EQUAL 是对列值进行比较的关系运算符
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))//期望列col2的值为3。
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)
	_, err := client.DeleteRow(deleteRowReq)
	if err !=nil{
		log.Fatal(err)
	}

}

