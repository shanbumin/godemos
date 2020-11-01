package install

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/moka-mrp/sword-core/samutils"
	"time"
)

//批量写入GSI2表
func  BatchWriteGSI2TableSample(client *tablestore.TableStoreClient, tableName string){

		batchWriteReq := &tablestore.BatchWriteRowRequest{}
		//批量接口一次只能插入200行
		for i := 1; i <= 200; i++ {
			putRowChange := new(tablestore.PutRowChange)
			//1.TableName
			putRowChange.TableName = tableName
			//2.主键
			putPk := new(tablestore.PrimaryKey)
			putPk.AddPrimaryKeyColumn("pk1",samutils.RandStringWordL(3))
			putPk.AddPrimaryKeyColumn("pk2",int64(i))
			putRowChange.PrimaryKey = putPk
			//3.属性列
			putRowChange.AddColumn("definedcol1",samutils.RandStringWordL(5))
			putRowChange.AddColumn("definedcol12",int64(time.Now().Unix()))
			//4.条件更新
			putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
            //追加改行
			batchWriteReq.AddRowChange(putRowChange)
		}
		response, err := client.BatchWriteRow(batchWriteReq)
		if err != nil {
			fmt.Println("batch request failed with:", response)
		} else {
			fmt.Println("batch write row finished")
		}


}


