package servers

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/constants"
)


//通过将IncludeBaseData参数设置为true，创建索引表后会开启数据表中存量数据的同步，然后可以通过索引表查询全部数据，
//同步时间跟数据量的大小有一定的关系。
func CreateGlobalIndexSample(client *tablestore.TableStoreClient, tableName string) {

	indexReq:=&tablestore.CreateIndexRequest{}


	//1.数据表名称
    indexReq.MainTableName=tableName
    //2.索引表的结构信息
	indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的DEFINED_COL_NAME_1列作为索引表的主键。索引表的索引列，索引列为数据表主键和预定义列的任意组合。
	indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。索引表的属性列，索引表属性列为数据表的预定义列的组合。
	indexMeta.IndexName = constants.GSI2Definedcol1Index
	indexReq.IndexMeta=indexMeta
    //3.索引表中是否包含数据表中已存在的数据
    indexReq.IncludeBaseData=true //包含存量数据


	resp, err := client.CreateIndex(indexReq)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create index finished", resp)
	}
}



//todo 由于系统会自动将未出现在索引列中的数据表主键补齐到索引表主键中，所以设置起始主键和结束主键时，需要同时设置索引表索引列和补齐的数据表主键。
//todo 与数据表的主键范围查询一模一样啦，就是表名换成对应的索引表名即可
func GetRangeFromIndex(client *tablestore.TableStoreClient, indexName string) {
	getRangeRequest := &tablestore.GetRangeRequest{}
	rangeRowQueryCriteria := &tablestore.RangeRowQueryCriteria{}
	//1.索引表表名
	rangeRowQueryCriteria.TableName = indexName
    //2.开始主键
	startPK := new(tablestore.PrimaryKey)
	startPK.AddPrimaryKeyColumnWithMinValue("pk1") //索引表的第一主键列。
	startPK.AddPrimaryKeyColumnWithMinValue("pk2") //索引表的第二主键列。
	startPK.AddPrimaryKeyColumnWithMinValue("pk3") //索引表的第三主键列。
	rangeRowQueryCriteria.StartPrimaryKey = startPK
	//3.结束主键
	endPK := new(tablestore.PrimaryKey)
	endPK.AddPrimaryKeyColumnWithMaxValue("pk1")
	endPK.AddPrimaryKeyColumnWithMaxValue("pk2")
	endPK.AddPrimaryKeyColumnWithMaxValue("pk3")
	rangeRowQueryCriteria.EndPrimaryKey = endPK
	//4.方向
	rangeRowQueryCriteria.Direction = tablestore.FORWARD
	//5.版本
	rangeRowQueryCriteria.MaxVersion = 1
	//6.返回个数
	rangeRowQueryCriteria.Limit = 10

	getRangeRequest.RangeRowQueryCriteria = rangeRowQueryCriteria
	getRangeResp, err := client.GetRange(getRangeRequest)
	for {
		if err != nil {
			fmt.Println("get range failed with error:", err)
		}
		if len(getRangeResp.Rows) > 0 {
			for _, row := range getRangeResp.Rows {
				//fmt.Println("range get row with key", row.PrimaryKey.PrimaryKeys[0].Value, row.PrimaryKey.PrimaryKeys[1].Value, row.PrimaryKey.PrimaryKeys[2].Value)
				fmt.Println(row)
			}
			if getRangeResp.NextStartPrimaryKey == nil {
				break
			} else {
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

//使用DeleteIndex接口删除数据表上指定的索引表。
func DeleteIndex(client *tablestore.TableStoreClient, tableName string,indexName string) {
	deleteIndex := &tablestore.DeleteIndexRequest{MainTableName: tableName, IndexName: indexName}

	resp, err := client.DeleteIndex(deleteIndex)
	if err != nil {
		fmt.Println("Failed to delete index:", err)
	} else {
		fmt.Println("drop index finished", resp)
	}

}