package sample

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
)


//PutRow接口用于新写入一行数据。如果该行已存在，则先删除原行数据（原行的所有列以及所有版本的数据），再写入新行数据。

// @param PutRowRequest    todo 执行PutRow操作所需参数的封装。
// @return PutRowResponse
func PutRowSample(client *tablestore.TableStoreClient, tableName string) {

	// 行的主键:设置的主键个数和类型必须和数据表的主键个数和类型一致。 当主键为自增列时，只需将相应主键指定为自增主键
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
	//putPk.AddPrimaryKeyColumnWithAutoIncrement("pk2")
	putPk.AddPrimaryKeyColumn("pk2", int64(2))
	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	//1.PutRowChange
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableName //数据表名称
	putRowChange.PrimaryKey = putPk //行的主键

	putRowChange.AddColumn("col1", "col1data3")
	putRowChange.AddColumn("col2", int64(3))
	putRowChange.AddColumn("col3", []byte("test3"))
	//使用条件更新，可以设置原行的存在性条件或者原行中某列的列值条件
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

	//2.PutRowRequest
	putRowRequest := new(tablestore.PutRowRequest)
	putRowRequest.PutRowChange = putRowChange

	//3.PutRow
	_, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("putrow failed with error:", err)
	} else {
		fmt.Println("putrow finished")
	}
}




// 返回表中的一行数据。
// @param GetRowRequest             执行GetRow操作所需参数的封装。
// @return  GetRowResponse          GetRow操作的响应内容。
func GetRowSample(client *tablestore.TableStoreClient, tableName string) {
    //1.PrimaryKey
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
	putPk.AddPrimaryKeyColumn("pk2",int64(2)) //todo 这里该怎么写呢
	//putPk.AddPrimaryKeyColumnWithAutoIncrement("pk2")
	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3"))

	//2.SingleRowQueryCriteria(查询条件)
	criteria := new(tablestore.SingleRowQueryCriteria)
	criteria.PrimaryKey = putPk

    //3.GetRowRequest
	getRowRequest := new(tablestore.GetRowRequest)
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = tableName
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1

	//4.GetRow
	getResp, err := client.GetRow(getRowRequest)
	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		colmap := getResp.GetColumnMap()
		fmt.Println("length is ", len(colmap.Columns))
		if len(colmap.Columns) >0{
			fmt.Println("get row col0 result is ", getResp.Columns[0].ColumnName, getResp.Columns[0].Value)
		}
	}




}

// @param UpdateRowRequest      执行UpdateRow操作所需参数的封装。
// @return UpdateRowResponse    UpdateRow操作的响应内容。

//todo 1.增加或更新数据时，需要设置属性名、属性值、属性类型（可选）、时间戳（可选）。
//todo 2.删除属性列特定版本的数据时，只需要设置属性名和时间戳。
//todo 3.删除属性列时，只需要设置属性名
//todo 4.说明 删除一行的全部属性列不等同于删除该行，如果需要删除该行，请使用DeleteRow操作。

func UpdateRowSample(client *tablestore.TableStoreClient, tableName string) {

    //1.PrimaryKey
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	//2.UpdateRowChange
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updateRowChange.PrimaryKey = updatePk
	updateRowChange.DeleteColumn("col1")
	updateRowChange.PutColumn("col2", int64(77))
	updateRowChange.PutColumn("col4", "newcol3")
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST) //使用条件更新：只有改行存在才更新成功

	//3.UpdateRowRequest
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowRequest.UpdateRowChange = updateRowChange

	//4.UpdateRow
	_, err := client.UpdateRow(updateRowRequest)

	if err != nil {
		fmt.Println("update failed with error:", err)
	} else {
		fmt.Println("update row finished")
	}
}

// @param DeleteRowRequest           执行DeleteRow操作所需参数的封装。
// @return DeleteRowResponse         DeleteRow操作的响应内容。
func DeleteRowSample(client *tablestore.TableStoreClient, tableName string) {

    //PrimaryKey
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	deletePk.AddPrimaryKeyColumn("pk2", int64(2))
	deletePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	//1.DeleteRowRequest
	deleteRowReq := new(tablestore.DeleteRowRequest)
	//2.DeleteRowChange
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	deleteRowReq.DeleteRowChange.TableName = tableName
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)
	_, err := client.DeleteRow(deleteRowReq)

	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		fmt.Println("delete row finished")
	}
}








































func UpdateRowWithIncrement(client *tablestore.TableStoreClient, tableName string) {
	fmt.Println("begin to update row")
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1increment")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	updateRowChange.PrimaryKey = updatePk

	updateRowChange.PutColumn("col2", int64(50))
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := client.UpdateRow(updateRowRequest)

	if err != nil {
		fmt.Println("update failed with error:", err)
		return
	} else {
		fmt.Println("update row finished")
	}

	updateRowRequest = new(tablestore.UpdateRowRequest)
	updateRowChange = new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updatePk = new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1increment")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	updateRowChange.PrimaryKey = updatePk

	updateRowChange.IncrementColumn("col2", int64(10))
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err = client.UpdateRow(updateRowRequest)
	if err != nil {
		fmt.Println("update failed with error:", err)
		return
	} else {
		fmt.Println("update row finished")
	}

	updateRowRequest = new(tablestore.UpdateRowRequest)
	updateRowChange = new(tablestore.UpdateRowChange)
	updateRowChange.TableName = tableName
	updatePk = new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1increment")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	updateRowChange.PrimaryKey = updatePk

	updateRowChange.IncrementColumn("col2", int64(30))
	updateRowChange.SetReturnIncrementValue()
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	updateRowChange.AppendIncrementColumnToReturn("col2")
	updateRowRequest.UpdateRowChange = updateRowChange

	resp, err := client.UpdateRow(updateRowRequest)
	if err != nil {
		fmt.Println("update failed with error:", err)
		return
	} else {
		fmt.Println("update row finished")
		fmt.Println(resp)
		fmt.Println(len(resp.Columns))
		fmt.Println(resp.Columns[0].ColumnName)
		fmt.Println(resp.Columns[0].Value)
		fmt.Println(resp.Columns[0].Timestamp)
	}
}

//func PutRowWithKeyAutoIncrementSample(client *tablestore.TableStoreClient) {
//	fmt.Println("begin to put row")
//	putRowRequest := new(tablestore.PutRowRequest)
//	putRowChange := new(tablestore.PutRowChange)
//	putRowChange.TableName = "incrementsampletable"
//	putPk := new(tablestore.PrimaryKey)
//	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
//	putPk.AddPrimaryKeyColumnWithAutoIncrement("pk2")
//	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
//	putRowChange.PrimaryKey = putPk
//	putRowChange.AddColumn("col1", "col1data1")
//	putRowChange.AddColumn("col2", int64(100))
//	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
//	putRowRequest.PutRowChange = putRowChange
//	_, err := client.PutRow(putRowRequest)
//
//	if err != nil {
//		fmt.Println("put row failed with error:", err)
//	} else {
//		fmt.Println("put row finished")
//	}
//}