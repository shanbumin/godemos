package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
	"otsdemo/bootstrap"
)

//查询基站002从时间1532574740开始的所有话单。
//但是查询不仅把base_station_number列作为条件，同时把start_time列作为查询条件，
//因此可以在base_station_number和start_time列上建立组合索引，索引表名称为sbm_call_record_gsi_base_station_number1，
//然后在索引表sbm_call_record_gsi_base_station_number1上进行查询。
func main() {
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	baseStationNumber:=2
	startTime:=1532574740
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
	endPK.AddPrimaryKeyColumnWithMaxValue("start_time")
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
				//todo 此索引表没有属性列，所以直接遍历主键数据即可
				//  for _,v:=range row.PrimaryKey.PrimaryKeys{
				//  	fmt.Println(v.ColumnName,v.Value)
				//  }
				fmt.Println(row.PrimaryKey.PrimaryKeys[0].ColumnName,row.PrimaryKey.PrimaryKeys[0].Value,
					row.PrimaryKey.PrimaryKeys[1].ColumnName,row.PrimaryKey.PrimaryKeys[1].Value,
					row.PrimaryKey.PrimaryKeys[2].ColumnName,row.PrimaryKey.PrimaryKeys[2].Value)
				//fmt.Println(row)
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
