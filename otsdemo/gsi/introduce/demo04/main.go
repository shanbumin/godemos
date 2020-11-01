package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
	"otsdemo/sdk/start"
)

//查询发生在基站003上时间从1532574861到1532584054的所有通话记录的通话时长。
//在该查询中不仅把base_station_number列和start_time列作为查询条件，而且要把duration列作为返回结果。
//您可以使用上一个查询中的索引表sbm_call_record_gsi_base_station_number1，查询索引表成功后反查数据表获取通话时长。
func main() {
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	baseStationNumber:=3
	startTime:=1532574861
	endTime:=1532584054
	//1.索引表表名
	rangeRowQueryCriteria.TableName = sample.GSI1BaseStationNumber1Index
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
	getRangeResp, err := start.Client.GetRange(getRangeRequest)
	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		if len(getRangeResp.Rows) > 0 {
			for _, row := range getRangeResp.Rows {
				curIndexPrimaryKey := row.PrimaryKey
				mainCalledNumber:=curIndexPrimaryKey.PrimaryKeys[2].Value
				callStartTime:=curIndexPrimaryKey.PrimaryKeys[1].Value
				//反查数据表-----
				getRowRequest := new(tablestore.GetRowRequest)
				criteria := new(tablestore.SingleRowQueryCriteria)
				criteria.TableName = sample.GSI1Table
				putPk := new(tablestore.PrimaryKey)
				putPk.AddPrimaryKeyColumn("cell_number",mainCalledNumber)
				putPk.AddPrimaryKeyColumn("start_time",callStartTime)
				criteria.PrimaryKey = putPk
				criteria.MaxVersion=1
				criteria.ColumnsToGet=[]string{"duration"}
				getRowRequest.SingleRowQueryCriteria = criteria
				getResp, err := start.Client.GetRow(getRowRequest)
				if err != nil {
					fmt.Println("getrow failed with error:", err)
				} else {
					colmap := getResp.GetColumnMap().Columns
					duration:=colmap["duration"][0].Value//值是多版本的，所以取第一个
					fmt.Println(duration)
				}
				//-----

			}
			if getRangeResp.NextStartPrimaryKey == nil {
				break
			} else {
				getRangeRequest.RangeRowQueryCriteria.StartPrimaryKey = getRangeResp.NextStartPrimaryKey
				getRangeResp, err = start.Client.GetRange(getRangeRequest)
			}
		} else {
			break
		}

		fmt.Println("continue to query rows")
	}
	fmt.Println("putrow finished")
}
