package servers

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/moka-mrp/sword-core/samutils"
	"time"
)


//PutRow接口用于新写入一行数据。
//如果该行已存在，则先删除原行数据（原行的所有列以及所有版本的数据），再写入新行数据。

func PutRowSample(client *tablestore.TableStoreClient, tableName string) {

	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)

	//1.数据表名称。
	putRowChange.TableName = tableName
	//2.行的主键。
	//设置的主键个数和类型必须和数据表的主键个数和类型一致。
	//当主键为自增列时，只需将相应主键指定为自增主键。
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("appid", "wx32d84d8f87a72033")
	putPk.AddPrimaryKeyColumn("openid","oVqItwN1CTLFCVd_KYtkeT8uUD_8")
	putRowChange.PrimaryKey = putPk
	//3.行的属性列Columns
	putRowChange.AddColumn("name","judy") //name唯一
	putRowChange.AddColumn("age",int64(28))
	putRowChange.AddColumn("salary",float64(3000.28))
	putRowChange.AddColumn("married",true)
	putRowChange.AddColumn("desc",[]byte(samutils.RandStringWordL(5)))
	putRowChange.AddColumn("created_at",int64(time.Now().Unix()))
	putRowChange.AddColumn("updated_at",int64(time.Now().Unix()))
	//4.使用条件更新，可以设置原行的存在性条件或者原行中某列的列值条件
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)


	putRowRequest.PutRowChange = putRowChange
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
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	//1.表名
	criteria.TableName = tableName
	//2.行的主键
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("appid", "wx32d84d8f87a72033")
	putPk.AddPrimaryKeyColumn("openid","oVqItwN1CTLFCVd_KYtkeT8uUD_8")
	criteria.PrimaryKey = putPk
	//3.最多读取的版本数。
	criteria.MaxVersion=1

	//GetRow
	getRowRequest.SingleRowQueryCriteria = criteria
	getResp, err := client.GetRow(getRowRequest)
	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		colmap := getResp.GetColumnMap().Columns
		for k,v:= range colmap{
           fmt.Println(k,v[0].Value) //值是多版本的，所以取第一个
		}
		//fmt.Println("length is ", len(colmap.Columns))
		//if len(colmap.Columns) >0{
		//	fmt.Println("get row col0 result is ", getResp.Columns[0].ColumnName, getResp.Columns[0].Value)
		//}
	}




}




func UpdateRowSample(client *tablestore.TableStoreClient, tableName string) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)

	//1.数据表名称。
	updateRowChange.TableName = tableName
	//2.行的主键
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("appid", "wx32d84d8f87a72033")
	updatePk.AddPrimaryKeyColumn("openid","oVqItwN1CTLFCVd_KYtkeT8uUD_8")

	updateRowChange.PrimaryKey = updatePk
	//3.行的属性列Columns
	//todo a.增加或更新数据时，需要设置属性名、属性值、属性类型（可选）、时间戳（可选）。
	//todo b.删除属性列特定版本的数据时，只需要设置属性名和时间戳。
	//todo c.删除属性列时，只需要设置属性名
	//todo d.说明 删除一行的全部属性列不等同于删除该行，如果需要删除该行，请使用DeleteRow操作。
	updateRowChange.DeleteColumn("desc") //删除
	updateRowChange.PutColumn("age", int64(29)) //修改
	updateRowChange.PutColumn("sex", "female") //故意新增一行
	//4.使用条件更新
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)



	updateRowRequest.UpdateRowChange = updateRowChange
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

	deleteRowReq := new(tablestore.DeleteRowRequest)
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)

	//1.数据表名称。
	deleteRowReq.DeleteRowChange.TableName = tableName

	//2.行的主键
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("appid", "wx32d84d8f87a72033")
	deletePk.AddPrimaryKeyColumn("openid","oVqItwN1CTLFCVd_KYtkeT8uUD_8")
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk

	//3.条件更新
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	clCondition1.FilterIfMissing=true //列不存在的时候不允许通过，否则删除的行不存在这个字段也删除成功了
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)



	_, err := client.DeleteRow(deleteRowReq)
	if err != nil {
		fmt.Println("delete failed with error:", err)
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