package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)

//查询号码234567的所有主叫话单。

//表格存储的模型是对所有行按照主键进行排序，并且提供顺序扫描（getRange）接口，
//所以只需要在调用getRange接口时，将cell_number列（主叫号码）的最大及最小值均设置为234567，
//start_time列（通话发生时间）的最小值设置为0，最大值设置为INT_MAX，对数据表进行扫描即可。

//todo 解决手法是 在数据表中进行主键的范围查询:多行数据操作-范围读（GetRange）
func main()  {

	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	cellNumber:=234567
	//1.表名
	rangeRowQueryCriteria.TableName = constants.GSI1Table
    //2.构建开始主键
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumn("cell_number",int64(cellNumber))
	startPK.AddPrimaryKeyColumn("start_time",int64(0))
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	//3.构建结束主键
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumn("cell_number",int64(cellNumber))
	endPK.AddPrimaryKeyColumnWithMaxValue("start_time")//这里等价给该值赋予了INT_MAX
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	//4.读取方向。
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	//5.最多读取的版本数。
	rangeRowQueryCriteria.MaxVersion = 1
	//6.数据的最大返回行数，此值必须大于 0。 todo 注意是返回行数,读懂它，就能做好分页
	rangeRowQueryCriteria.Limit = 2

	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
	getRangeResp, err := bootstrap.Client.GetRange(getRangeRequest)

	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		for _, row := range getRangeResp.Rows {
			fmt.Println(row)
		}
		if getRangeResp.NextStartPrimaryKey == nil {
			break
		} else { // 若nextStartPrimaryKey不为nil, 则继续读取。
			getRangeRequest.RangeRowQueryCriteria.StartPrimaryKey = getRangeResp.NextStartPrimaryKey
			getRangeResp, err = bootstrap.Client.GetRange(getRangeRequest)
		}
		fmt.Println("continue to query rows")
	}
	fmt.Println("range get row finished")


}
