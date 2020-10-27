package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"log"
	"otsdemo/sample"
)
/*

todo 条件更新:只有满足条件时，才能对数据表中的数据进行更新；当不满足条件时，更新失败。
todo 在通过PutRow、UpdateRow、DeleteRow或BatchWriteRow接口更新数据时，可以使用条件更新检查行存在性条件和列条件，只有满足条件时才能更新成功。
todo 条件更新包括行存在性条件和列条件。

行存在性条件：包括IGNORE、EXPECT_EXIST和EXPECT_NOT_EXIST，分别代表忽略、期望存在和期望不存在。
对数据表进行更改操作时，系统会先检查行存在性条件，如果不满足行存在性条件，则更改失败并给用户报错。

列条件：包括SingleColumnCondition和CompositeColumnCondition，是基于某一列或者某些列的列值进行条件判断。
SingleColumnCondition支持一列（可以是主键列）和一个常量比较。不支持两列或者两个常量相比较。
CompositeColumnCondition的内节点为逻辑运算，子条件可以是SingleColumnCondition或CompositeColumnCondition。


todo 条件更新可以实现乐观锁功能，即在更新某行时，先获取某列的值，假设为列A，值为1，然后设置条件列A＝1，更新行使列A＝2。如果更新失败，表示有其他客户端已成功更新该行。



updateRowChange.SetCondition( tablestore.RowExistenceExpectation_IGNORE)    :忽略
updateRowChange.SetCondition( tablestore.RowExistenceExpectation_EXPECT_EXIST)   :期望存在
updateRowChange.SetCondition( tablestore.RowExistenceExpectation_EXPECT_NOT_EXIST)   :期望不存在

*/

func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.根据指定主键更新一行，如果指定的行存在，则更新成功，否则更新失败。
	tableName:="t11"
   //a.PrimaryKey
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))

	//b.UpdateRowChange
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.DeleteColumn("col1")            //删除col1列。
	updateRowChange.PutColumn("col2", int64(77))    //新增col2列，值为77。
	updateRowChange.PutColumn("col4", "newcol3")    //新增col4列，值为"newcol3"。
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST) //期望指定行存在。

	//c.UpdateRowRequest
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := client.UpdateRow(updateRowRequest)
	if err !=nil{
		log.Fatal(err)
	}





}
