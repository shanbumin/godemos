package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)


//为了提高查询效率，可以在base_station_number列和start_time列上建立组合索引，并把duration列作为索引表的属性列，
//索引表名称为sbm_call_record_gsi_base_station_number2，然后在索引表sbm_call_record_gsi_base_station_number2上进行查询。

func main() {
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	baseStationNumber:=3
	startTime:=1532574861
	endTime:=1532584054
	//1.索引表表名
	rangeRowQueryCriteria.TableName = sample.GSI1BaseStationNumber2Index
	//2.开始主键
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumn("base_station_number",int64(baseStationNumber))
	startPK.AddPrimaryKeyColumn("start_time",int64(startTime))
	startPK.AddPrimaryKeyColumnWithMinValue("cell_number")
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	//3.结束主键
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumn("base_station_number",int64(baseStationNumber))
	endPK.AddPrimaryKeyColumn("start_time",int64(endTime))
	endPK.AddPrimaryKeyColumnWithMaxValue("cell_number")
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	//4.方向
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	//5.版本
	rangeRowQueryCriteria.MaxVersion = 1
	//6.返回个数
	rangeRowQueryCriteria.Limit = 10

	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
	getRangeResp, err := bootstrap.Client.GetRange(getRangeRequest)
	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		if len(getRangeResp.Rows) > 0 {
			for _, row := range getRangeResp.Rows {
				curIndexPrimaryKey := row.PrimaryKey
				curColumn:=row.Columns

				fmt.Println(curIndexPrimaryKey.PrimaryKeys[0].ColumnName,curIndexPrimaryKey.PrimaryKeys[0].Value,
					curIndexPrimaryKey.PrimaryKeys[1].ColumnName,curIndexPrimaryKey.PrimaryKeys[1].Value,
					curIndexPrimaryKey.PrimaryKeys[2].ColumnName,curIndexPrimaryKey.PrimaryKeys[2].Value,
					curColumn[0].ColumnName,curColumn[0].Value)
				//-----

			}
			if getRangeResp.NextStartPrimaryKey == nil {
				break
			} else {
				getRangeRequest.RangeRowQueryCriteria.StartPrimaryKey = getRangeResp.NextStartPrimaryKey
				getRangeResp, err = bootstrap.Client.GetRange(getRangeRequest)
			}
		} else {
			break
		}

		fmt.Println("continue to query rows")
	}
	fmt.Println("putrow finished")
}
