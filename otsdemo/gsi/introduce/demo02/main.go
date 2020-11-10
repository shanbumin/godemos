package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)

//查询号码123456的被叫话单。
//表格存储的模型是对所有行按照主键进行排序，由于被叫号码存在于数据表的预定义列中，所以无法进行快速查询。
//因此可以在被叫号码索引表sbm_call_record_gsi_called_number上进行查询，
//由于索引表sbm_call_record_gsi_called_number是按照被叫号码作为主键，可以直接调用getRange接口扫描索引表得到结果。
func main()  {

	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	calledNumber:=123456
	//1.索引表表名
	rangeRowQueryCriteria.TableName = constants.GSI1CalledNumberIndex
	//2.开始主键
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumn("called_number",int64(calledNumber))
	startPK.AddPrimaryKeyColumnWithMinValue("cell_number") //索引表的第二主键列。
	startPK.AddPrimaryKeyColumnWithMinValue("start_time") //索引表的第三主键列。
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	//3.结束主键
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumn("called_number",int64(calledNumber))
	endPK.AddPrimaryKeyColumnWithMaxValue("cell_number")
	endPK.AddPrimaryKeyColumnWithMaxValue("start_time")
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
