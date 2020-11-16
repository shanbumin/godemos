package install

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
)

//var gsi1Data=map[int]map[string]interface{}{
//}
//数组定义并初始化
var gsi1Data=[...]map[string]interface{}{
	{"cell_number":123456,"start_time":1532574644,"called_number":654321,"duration":60,"base_station_number":1},
	{"cell_number":234567,"start_time":1532574714,"called_number":765432,"duration":10,"base_station_number":1},
	{"cell_number":234567,"start_time":1532574734,"called_number":123456,"duration":20,"base_station_number":3},
	{"cell_number":345678,"start_time":1532574795,"called_number":123456,"duration":5,"base_station_number":2},
	{"cell_number":345678,"start_time":1532574861,"called_number":123456,"duration":100,"base_station_number":2},
	{"cell_number":456789,"start_time":1532584054,"called_number":345678,"duration":200,"base_station_number":3},

	{"cell_number":234567,"start_time":1601524328,"called_number":345678,"duration":10,"base_station_number":3},
	{"cell_number":234567,"start_time":1601610728,"called_number":345678,"duration":100,"base_station_number":3},
	{"cell_number":234567,"start_time":1604202728,"called_number":345678,"duration":10,"base_station_number":3},
}


//批量写入GSI1表
func  BatchWriteGSI1TableSample(client *tablestore.TableStoreClient, tableName string){
	batchWriteReq := &tablestore.BatchWriteRowRequest{}
	for  _,v:=range  gsi1Data{
		//----
		putRowChange := new(tablestore.PutRowChange)
		putRowChange.TableName = tableName
		putPk := new(tablestore.PrimaryKey)
		putPk.AddPrimaryKeyColumn("cell_number",int64(v["cell_number"].(int))) //todo 这里我们已经知道它是int了，所以直接转，不用类型断言了
		putPk.AddPrimaryKeyColumn("start_time", int64(v["start_time"].(int)))
		putRowChange.PrimaryKey = putPk
		putRowChange.AddColumn("called_number",int64(v["called_number"].(int)))
		putRowChange.AddColumn("duration",int64(v["duration"].(int)))
		putRowChange.AddColumn("base_station_number",int64(v["base_station_number"].(int)))
		putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
		batchWriteReq.AddRowChange(putRowChange)
		//----
	}
	response, err := client.BatchWriteRow(batchWriteReq)
	if err != nil {
		fmt.Println("batch request failed with:", response)
	} else {
		fmt.Println("batch write row finished")
	}
}

