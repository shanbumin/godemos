package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/moka-mrp/sword-core/samutils"
	"otsdemo/sample"
	"otsdemo/bootstrap"
	"strconv"
	"time"
)



func  RandBool()bool{
	res:=samutils.MtRand(0,2)
	if res==0{
		return false
	}
	return true
}
func main() {





	for  i:=1;i<=1;i++{


		batchWriteReq := &tablestore.BatchWriteRowRequest{}
		//批量接口一次只能插入200行
		for j := 1; j <= 200; j++ {
			putRowChange := new(tablestore.PutRowChange)
			//1.TableName
			putRowChange.TableName = sample.DemoTableName
			//2.主键
			name:="sam"+strconv.Itoa(i*j)
			putPk := new(tablestore.PrimaryKey)
			putPk.AddPrimaryKeyColumn("_id",samutils.Md5(name))
			putRowChange.PrimaryKey = putPk
			//3.属性列
			putRowChange.AddColumn("name", name) //name唯一
			putRowChange.AddColumn("age",int64(j))
			putRowChange.AddColumn("salary",float64(j*100))
			putRowChange.AddColumn("married",RandBool())
			putRowChange.AddColumn("desc",[]byte(samutils.RandStringWordL(5)))
			putRowChange.AddColumn("created_at",int64(time.Now().Unix()))
			putRowChange.AddColumn("updated_at",int64(time.Now().Unix()))
			//4.条件更新
			putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

			batchWriteReq.AddRowChange(putRowChange)
		}
		response, err := bootstrap.Client.BatchWriteRow(batchWriteReq)
		if err != nil {
			fmt.Println("batch request failed with:", response)
		} else {
			fmt.Println("batch write row finished",i)
		}



	}




}
