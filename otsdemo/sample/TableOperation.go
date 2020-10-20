package sample
import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

//创建表1
func CreateTableSample001(client *tablestore.TableStoreClient, tableName string) {

	//1.tableMeta (创建主键列的schema)
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk0", tablestore.PrimaryKeyType_INTEGER) //第一个PK列为整数，名称是pk0，此列同时也是分区键。
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING) //第二个PK列为整数，名称是pk1。
	//2.tableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1 //数据生命周期
	tableOption.MaxVersion = 3   //最大版本数

	//3.reservedThroughput(保留吞吐量)
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0 //预留读吞吐量
	reservedThroughput.Writecap = 0  //预留写吞吐量


	//4.createtableRequest
	createtableRequest := new(tablestore.CreateTableRequest)
	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput

	//5.create
	_, err := client.CreateTable(createtableRequest)
	if (err != nil) {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}

}

//创建表2
func CreateTableWithGlobalIndexSample(client *tablestore.TableStoreClient, tableName string) {

	//1.tableMeta (创建主键列的schema)
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddDefinedColumn("definedcol1", tablestore.DefinedColumn_STRING) //添加预定义列1
	tableMeta.AddDefinedColumn("definedcol2", tablestore.DefinedColumn_INTEGER) //添加预定义列2

	indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的definedcol1列作为索引表的主键。
	indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。
	indexMeta.IndexName = "indexSample"
	//2.tableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1 //数据的过期时间，单位为秒，-1表示永不过期。带索引表的数据表数据生命周期必须设置为-1。
	tableOption.MaxVersion = 1 //保存的最大版本数，1表示每列上最多保存一个版本即保存最新的版本。带索引表的数据表最大版本数必须设置为1。
	//3.reservedThroughput(保留吞吐量)
	reservedThroughput := new(tablestore.ReservedThroughput)

	//4.createtableRequest
	createTableRequest := new(tablestore.CreateTableRequest)
	createTableRequest.TableMeta = tableMeta
	createTableRequest.TableOption = tableOption
	createTableRequest.ReservedThroughput = reservedThroughput
	createTableRequest.AddIndexMeta(indexMeta)
	//5.create
	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}
}


//----------------------------------------------------------------------------------------------------------------------

func UpdateTableSample(client *tablestore.TableStoreClient, tableName string) {

	//1.UpdateTableRequest
	updateTableReq := new(tablestore.UpdateTableRequest)
	updateTableReq.TableName = tableName
	updateTableReq.TableOption = new(tablestore.TableOption)
	updateTableReq.TableOption.TimeToAlive = -1
	updateTableReq.TableOption.MaxVersion = 5

	//2.update
	_, err := client.UpdateTable(updateTableReq)
	if err != nil {
		fmt.Println("failed to update table with error:", err)
	} else {
		fmt.Println("update finished")
	}
}

//----------------------------------------------------------------------------------------------------------------------

//列出所有的表，如果操作成功，将返回所有表的名称。
//ListTable() (*ListTableResponse, error)


func ListTableSample(client *tablestore.TableStoreClient) {
	listtables, err := client.ListTable()
	if err != nil {
		fmt.Println("Failed to list table")
	} else {
		fmt.Println("List table result is")
		for _, table := range listtables.TableNames {
			fmt.Println("TableName: ", table)
		}
	}
}








func CreateTableSample(client *tablestore.TableStoreClient, tableName string) {


	fmt.Println("Begin to create table:", tableName)
	createtableRequest := new(tablestore.CreateTableRequest)
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0


	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput


	_, err := client.CreateTable(createtableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}


}






func CreateTableKeyAutoIncrementSample(client *tablestore.TableStoreClient) {
	createtableRequest := new(tablestore.CreateTableRequest)
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = "incrementsampletable"
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumnOption("pk2", tablestore.PrimaryKeyType_INTEGER, tablestore.AUTO_INCREMENT)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput
	client.CreateTable(createtableRequest)
}
func DeleteTableSample(client *tablestore.TableStoreClient) {
	fmt.Println("Begin to delete table")
	tableName := "tabletodeletesample"
	CreateTableSample(client, tableName)
	fmt.Println("Begin to delete table:", tableName)
	deleteReq := new(tablestore.DeleteTableRequest)
	deleteReq.TableName = tableName
	_, err := client.DeleteTable(deleteReq)
	if err != nil {
		fmt.Println("Failed to delete table with error:", err)
	} else {
		fmt.Println("Delete table finished")
	}
}


func DescribeTableSample(client *tablestore.TableStoreClient, tableName string) {
	fmt.Println("DescribeTableSample started")
	describeTableReq := new(tablestore.DescribeTableRequest)
	describeTableReq.TableName = tableName
	describ, err := client.DescribeTable(describeTableReq)
	if err != nil {
		fmt.Println("failed to update table with error:", err)
	} else {
		fmt.Println("DescribeTableSample finished. Table meta:", describ.TableOption.MaxVersion, describ.TableOption.TimeToAlive)
	}
}
func ComputeSplitPointsBySize(client *tablestore.TableStoreClient, tableName string) {
	req := &tablestore.ComputeSplitPointsBySizeRequest{TableName: tableName, SplitSize: int64(1)}
	va, err := client.ComputeSplitPointsBySize(req)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range va.Splits {
		fmt.Println(val.Location)
		fmt.Println(*val.LowerBound)
		fmt.Println(*val.UpperBound)
	}
	return
}

