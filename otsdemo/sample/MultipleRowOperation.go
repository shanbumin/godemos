package sample

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"math/rand"
	"time"
)


// 对多个数据表中的多行数据进行增加、删除或者更新操作。
// @param BatchWriteRowRequest             执行BatchWriteRow操作所需参数的封装。
// @return  BatchWriteRowResponse          BatchWriteRow操作的响应内容。
func BatchWriteRowSample(client *tablestore.TableStoreClient, tableName string) {
	batchWriteReq := &tablestore.BatchWriteRowRequest{}

	for i := 0; i < 100; i++ {
		//PrimaryKey
		putPk := new(tablestore.PrimaryKey)
		putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
		putPk.AddPrimaryKeyColumn("pk2", int64(i))
		putPk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
		//PutRowChange
		putRowChange := new(tablestore.PutRowChange)
		putRowChange.TableName = tableName
		putRowChange.PrimaryKey = putPk
		putRowChange.AddColumn("col1", "fixvalue")
		putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
		//AddRowChange
		batchWriteReq.AddRowChange(putRowChange)
	}

	//BatchWriteRow
	response, err := client.BatchWriteRow(batchWriteReq)
	if err != nil {
		fmt.Println("batch request failed with:", response)
	} else {
		// todo check all succeed
		fmt.Println("batch write row finished")
	}
}


//返回数据表（Table）中的多行数据。
// @param BatchGetRowRequest             执行BatchGetRow操作所需参数的封装。
// @return  BatchGetRowResponse          BatchGetRow操作的响应内容。
func BatchGetRowSample(client *tablestore.TableStoreClient, tableName string) {

	batchGetReq := &tablestore.BatchGetRowRequest{}
	mqCriteria := &tablestore.MultiRowQueryCriteria{}

	for i := 0; i < 20; i++ {
		pkToGet := new(tablestore.PrimaryKey)
		pkToGet.AddPrimaryKeyColumn("pk1", "pk1value1")
		pkToGet.AddPrimaryKeyColumn("pk2", int64(i))
		pkToGet.AddPrimaryKeyColumn("pk3", []byte("pk3"))
		mqCriteria.AddRow(pkToGet)
	}

	pkToGet2 := new(tablestore.PrimaryKey)
	pkToGet2.AddPrimaryKeyColumn("pk1", "pk1value2")
	pkToGet2.AddPrimaryKeyColumn("pk2", int64(300))
	pkToGet2.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	mqCriteria.AddColumnToGet("col1")
	mqCriteria.AddRow(pkToGet2)

	mqCriteria.MaxVersion = 1
	mqCriteria.TableName = tableName
	batchGetReq.MultiRowQueryCriteria = append(batchGetReq.MultiRowQueryCriteria, mqCriteria)

	/*
	condition := tablestore.NewSingleColumnCondition("col1", tablestore.CT_GREATER_THAN, int64(0))
	mqCriteria.Filter = condition
	*/

	batchGetResponse, err := client.BatchGetRow(batchGetReq)

	if err != nil {
		fmt.Println("batachget failed with error:", err)
	} else {
		for _, row := range batchGetResponse.TableToRowsResult[mqCriteria.TableName] {
			if row.PrimaryKey.PrimaryKeys != nil {
				fmt.Println("get row with key", row.PrimaryKey.PrimaryKeys[0].Value,
					row.PrimaryKey.PrimaryKeys[1].Value,
					row.PrimaryKey.PrimaryKeys[2].Value)
			} else {
				fmt.Println("this row is not exist")
			}
		}
		fmt.Println("batchget finished")
	}
}



// 从表中查询一个范围内的多行数据。
// @param GetRangeRequest            执行GetRange操作所需参数的封装。
// @return GetRangeResponse          GetRange操作的响应内容。
//todo 数据表中的行按主键从小到大排序，读取范围是一个左闭右开的区间，正序读取时，返回的是大于等于起始主键且小于结束主键的所有的行。
//todo 同一表中有两个主键A和B，A<B。如正序读取[A, B)，则按从A至B的顺序返回主键大于等于A、小于B的行；逆序读取[B, A)，则按从B至A的顺序返回大于A、小于等于B的数据。
func GetRangeSample(client *tablestore.TableStoreClient, tableName string) {
	fmt.Println("Begin to scan the table")


    //1.StartPrimaryKey表示起始主键，如果该行存在，则返回结果中一定会包含此行。
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumnWithMinValue("pk1")
	startPK.AddPrimaryKeyColumnWithMinValue("pk2")
	startPK.AddPrimaryKeyColumnWithMinValue("pk3")
	//2.EndPrimaryKey表示结束主键，无论该行是否存在，返回结果中都不会包含此行。
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumnWithMaxValue("pk1")
	endPK.AddPrimaryKeyColumnWithMaxValue("pk2")
	endPK.AddPrimaryKeyColumnWithMaxValue("pk3")
	//3.rangeRowQueryCriteria
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	rangeRowQueryCriteria.TableName = tableName
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	rangeRowQueryCriteria.MaxVersion = 1
	rangeRowQueryCriteria.Limit = 10

	//4.getRangeRequest
	getRangeRequest := &tablestore.GetRangeRequest{}
	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria

	//5.GetRange
	getRangeResp, err := client.GetRange(getRangeRequest)
	//fmt.Println("get range result is ", getRangeResp)
	//6.遍历结果
	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		if len(getRangeResp.Rows) > 0 {
			for _, row := range getRangeResp.Rows {
				fmt.Println("range get row with key", row.PrimaryKey.PrimaryKeys[0].Value,
					row.PrimaryKey.PrimaryKeys[1].Value,
					row.PrimaryKey.PrimaryKeys[2].Value)
			}
			if getRangeResp.NextStartPrimaryKey == nil {
				break
			} else {
				fmt.Println("next pk is :", getRangeResp.NextStartPrimaryKey.PrimaryKeys[0].Value,
					getRangeResp.NextStartPrimaryKey.PrimaryKeys[1].Value,
					getRangeResp.NextStartPrimaryKey.PrimaryKeys[2].Value)

				getRangeRequest.RangeRowQueryCriteria.StartPrimaryKey = getRangeResp.NextStartPrimaryKey
				getRangeResp, err = client.GetRange(getRangeRequest)
			}
		} else {
			break
		}

		fmt.Println("continue to query rows")
	}
	fmt.Println("putrow finished")

}









































































var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randStringRunes(random *rand.Rand, n int) string {
	//random := rand.New(rand.NewSource(time.Now().Unix()))

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[random.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetRangeWithRegxFilterSample(client *tablestore.TableStoreClient, tableName string) {
	fmt.Println("Begin to create table:", tableName)
	createtableRequest := new(tablestore.CreateTableRequest)

	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput

	_, err := client.CreateTable(createtableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}

	fmt.Println("batch write row started")
	batchWriteReq := &tablestore.BatchWriteRowRequest{}
	random := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 100; i++ {
		putRowChange := new(tablestore.PutRowChange)
		putRowChange.TableName = tableName
		putPk := new(tablestore.PrimaryKey)
		putPk.AddPrimaryKeyColumn("pk1", random.Int63n(10000))
		putPk.AddPrimaryKeyColumn("pk2", random.Int63n(10000))

		putRowChange.PrimaryKey = putPk
		colKey1 := randStringRunes(random, 5)
		colKey2 := randStringRunes(random, 5)
		val1 := "t1:" + colKey1 + "," + "t2:" + randStringRunes(random, 1) + "," + "t3:-" + randStringRunes(random, 1) + "," + "t4:" + randStringRunes(random, 1) + "." + randStringRunes(random, 1) + "," + "t5:dummy";
		val2 := "c1:" + colKey2 + "," + "c2:" + randStringRunes(random, 1) + "," + "c3:-" + randStringRunes(random, 1) + "," + "c4:" + randStringRunes(random, 1) + "." + randStringRunes(random, 1) + "," + "c5:dummy";
		putRowChange.AddColumn("col1", val1)
		putRowChange.AddColumn("col2", val2)
		putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
		batchWriteReq.AddRowChange(putRowChange)
	}

	response, err := client.BatchWriteRow(batchWriteReq)
	if err != nil {
		fmt.Println("batch request failed with:", response)
	} else {
		// todo check all succeed
		fmt.Println("batch write row finished")
	}

	fmt.Println("begin to range query with filter")

	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	rangeRowQueryCriteria.TableName = tableName

	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumnWithMinValue("pk1")
	startPK.AddPrimaryKeyColumnWithMinValue("pk2")
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumnWithMaxValue("pk1")
	endPK.AddPrimaryKeyColumnWithMaxValue("pk2")

	rangeRowQueryCriteria.StartPrimaryKey = startPK
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	rangeRowQueryCriteria.MaxVersion = 1
	rangeRowQueryCriteria.Limit = 1000
	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
	filter := tablestore.NewCompositeColumnCondition(tablestore.LogicalOperator(tablestore.LO_AND))
	regexFule1 := tablestore.NewValueTransferRule("t1:([a-z]+),", tablestore.Variant_STRING)
	filter1 := tablestore.NewSingleColumnValueRegexFilter("col1", tablestore.ComparatorType(tablestore.CT_GREATER_EQUAL), regexFule1, "d")
	regexFule2 := tablestore.NewValueTransferRule("t1:([a-z]+),", tablestore.Variant_STRING)
	filter2 := tablestore.NewSingleColumnValueRegexFilter("col1", tablestore.ComparatorType(tablestore.CT_LESS_EQUAL), regexFule2, "m")
	filter.AddFilter(filter1)
	filter.AddFilter(filter2)
	//getRangeRequest.RangeRowQueryCriteria.Filter = filter
	getRangeResp, err := client.GetRange(getRangeRequest)
	fmt.Println(err)
	//fmt.Println("get range result is ", getRangeResp.Rows)
	fmt.Println(getRangeResp.NextStartPrimaryKey)
	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		if len(getRangeResp.Rows) > 0 {
			for _, row := range getRangeResp.Rows {
				fmt.Println("range get row with key", row.PrimaryKey.PrimaryKeys[0].Value, row.PrimaryKey.PrimaryKeys[1].Value, row.Columns[0].ColumnName,row.Columns[0].Value)
			}
			if getRangeResp.NextStartPrimaryKey == nil {
				break
			} else {
				fmt.Println("next pk is :", getRangeResp.NextStartPrimaryKey.PrimaryKeys[0].Value, getRangeResp.NextStartPrimaryKey.PrimaryKeys[1].Value, getRangeResp.NextStartPrimaryKey.PrimaryKeys[2].Value)
				getRangeRequest.RangeRowQueryCriteria.StartPrimaryKey = getRangeResp.NextStartPrimaryKey
				getRangeResp, err = client.GetRange(getRangeRequest)
			}
		} else {
			break
		}

		fmt.Println("continue to query rows")
	}
	fmt.Println("putrow finished")
}
