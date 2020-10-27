package sample

import "github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
import "fmt"


func CreateGlobalIndexSample(client *tablestore.TableStoreClient, tableName string) {

	indexReq:=&tablestore.CreateIndexRequest{}


	//1.数据表名称
    indexReq.MainTableName=tableName
    //2.索引表的结构信息
	indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的DEFINED_COL_NAME_1列作为索引表的主键。索引表的索引列，索引列为数据表主键和预定义列的任意组合。
	indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。索引表的属性列，索引表属性列为数据表的预定义列的组合。
	indexMeta.IndexName = "indexSample"//索引表名称。
	indexReq.IndexMeta=indexMeta
    //3.索引表中是否包含数据表中已存在的数据
    indexReq.IncludeBaseData=true //包含存量数据



	/**
	  通过将IncludeBaseData参数设置为true，创建索引表后会开启数据表中存量数据的同步，然后可以通过索引表查询全部数据，
	  同步时间跟数据量的大小有一定的关系。
	*/
	resp, err := client.CreateIndex(indexReq)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create index finished", resp)
	}
}



