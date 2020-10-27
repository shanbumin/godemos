package sample
import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
)

//创建数据表
func CreateTableSample(client *tablestore.TableStoreClient, tableName string) {
	createtableRequest := new(tablestore.CreateTableRequest)

     //1.数据表的结构信息 TableMeta
	tableMeta := new(tablestore.TableMeta)
	//1.1 数据表名称
	tableMeta.TableName = tableName
	//1.2 数据表的主键定义
	// 第一列主键作为分区键。分区键相同的数据会存放在同一个分区内，所以相同分区键下最好不要超过10 GB以上数据，否则会导致单分区过大，无法分裂。
	// 另外，数据的读/写访问最好在不同的分区键上均匀分布，有利于负载均衡。
	tableMeta.AddPrimaryKeyColumn("pk0", tablestore.PrimaryKeyType_INTEGER) //第一个PK列为整数，名称是pk0，此列同时也是分区键。
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING) //第二个PK列为字符串，名称是pk1。
	//1.3 预先定义一些非主键列以及其类型，可以作为索引表的属性列或索引列。 DefinedColumns
	//说明 属性列不需要定义。表格存储每行的数据列都可以不同，属性列的列名在写入时指定。
	createtableRequest.TableMeta = tableMeta


	//2.数据表的配置信息
	//2.1 数据生命周期，即数据的过期时间。
	// 数据生命周期至少为86400秒（一天）或-1（数据永不过期）
	// 如果需要使用索引，则数据生命周期必须设置为-1（数据永不过期）
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	//2.2 最大版本数 MaxVersion
	//属性列能够保留数据的最大版本个数。当属性列数据的版本个数超过设置的最大版本数时，系统会自动删除较早版本的数据。
	//说明 如果需要使用索引，则最大版本数必须设置为1。
	tableOption.MaxVersion = 3
	//2.3 有效版本偏差 DeviationCellVersionInSec
	//即写入数据的时间戳与系统当前时间的偏差允许最大值。只有当写入数据所有列的版本号与写入时时间的差值在数据有效版本偏差范围内，数据才能成功写入。
	//属性列的有效版本范围为[数据写入时间-有效版本偏差，数据写入时间+有效版本偏差)。
	//创建数据表时，如果未设置有效版本偏差，系统会使用默认值86400
	createtableRequest.TableOption = tableOption

	//3.为数据表配置预留读吞吐量或预留写吞吐量。reservedThroughput
	//容量型实例中的数据表的预留读/写吞吐量只能设置为0，不允许预留。
	//默认值为0，即完全按量计费。
	//单位为CU。
	//当预留读吞吐量或预留写吞吐量大于0时，表格存储会根据配置为数据表预留相应资源，且数据表创建成功后，将会立即按照预留吞吐量开始计费，超出预留的部分进行按量计费。
	//当预留读吞吐量或预留写吞吐量设置为0时，表格存储不会为数据表预留相应资源。
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0 //预留读吞吐量
	reservedThroughput.Writecap = 0  //预留写吞吐量
	createtableRequest.ReservedThroughput = reservedThroughput

	//4.索引表的结构信息 IndexMeta
	//todo 略




	_, err := client.CreateTable(createtableRequest)
	if (err != nil) {
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

//----------------------------------------------------------------------------------------------------------------------
//通过表名查询表描述信息。
//DescribeTable(request *DescribeTableRequest) (*DescribeTableResponse, error)

func DescribeTableSample(client *tablestore.TableStoreClient, tableName string) {


	//1.DescribeTableRequest
	describeTableReq := new(tablestore.DescribeTableRequest)
	describeTableReq.TableName = tableName

	//2.DescribeTable
	describ, err := client.DescribeTable(describeTableReq)

	fmt.Printf("%#v\r\n",describ)
	if err != nil {
		fmt.Println("failed to update table with error:", err)
	} else {
		fmt.Println("DescribeTableSample finished. Table meta:", describ.TableOption.MaxVersion, describ.TableOption.TimeToAlive)
	}
}

//--------------------------------------------------------------------------------

//为数据表预先定义一些非主键列以及其类型，可以作为索引表的属性列或索引列。包含如下设置：
//Name：预定义列名称。
//ColumnType：预定义列的数据类型。


//为数据表增加预定义列，预定义列分别为definedColumnName01（String类型）、definedColumnName02（INTEGER类型）、definedColumnName03（String类型）。
func AddDefinedColumn(client *tablestore.TableStoreClient, tableName string) {
    //1.AddDefinedColumnRequest)
	addDefinedColumnRequest := new(tablestore.AddDefinedColumnRequest)
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName01",tablestore.DefinedColumn_STRING)
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName02",tablestore.DefinedColumn_INTEGER)
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName03",tablestore.DefinedColumn_STRING)
	addDefinedColumnRequest.TableName = tableName

	//2.AddDefinedColumn
	_, err := client.AddDefinedColumn(addDefinedColumnRequest)
	if (err != nil) {
		fmt.Println("Failed to Add DefinedColumn with error:", err)
	} else {
		fmt.Println("Add DefinedColumn finished")
	}
}


//删除数据表的预定义列definedColumnName01和definedColumnName02。
func DeleteDefinedColumn(client *tablestore.TableStoreClient, tableName string){

	//1.DeleteDefinedColumnRequest
	deleteDefinedColumnRequest := new(tablestore.DeleteDefinedColumnRequest)
	deleteDefinedColumnRequest.DefinedColumns = []string{"definedColumnName01","definedColumnName02"}
	deleteDefinedColumnRequest.TableName = tableName

	//2.DeleteDefinedColumn
	_, err := client.DeleteDefinedColumn(deleteDefinedColumnRequest)
	if (err != nil) {
		fmt.Println("Failed to delete DefinedColumn with error:", err)
	} else {
		fmt.Println("Delete DefinedColumn finished")
	}
}

//--------------------------------------------------------------------------------------------------------------------
//删除表
func DeleteTableSample(client *tablestore.TableStoreClient,tableName string) {
    //1.DeleteTableRequest
	deleteReq := new(tablestore.DeleteTableRequest)
	deleteReq.TableName = tableName

	//2.DeleteTable
	_, err := client.DeleteTable(deleteReq)
	if err != nil {
		fmt.Println("Failed to delete table with error:", err)
	} else {
		fmt.Println("Delete table finished")
	}
}

//--------------------------------------------------------------------------------------------------------------------

//创建表时，只需将自增的主键属性设置为AUTO_INCREMENT。
func CreateTableKeyAutoIncrementSample(client *tablestore.TableStoreClient,tableName string) {
     //1.meta   数据表中包括三个主键：pk1，String类型；pk2，INTEGER类型，为自增列；pk3，Binary类型。
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumnOption("pk2", tablestore.PrimaryKeyType_INTEGER, tablestore.AUTO_INCREMENT)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	//2.option
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	//3.reserved
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0

	//4.request

	createtableRequest := new(tablestore.CreateTableRequest)
	createtableRequest.TableMeta = tableMeta
	createtableRequest.TableOption = tableOption
	createtableRequest.ReservedThroughput = reservedThroughput

	//5.create
	client.CreateTable(createtableRequest)
}


func PutRowWithKeyAutoIncrementSample(client *tablestore.TableStoreClient,tableName string) {

	//1.设置主键，必须按照创建数据表时的顺序添加主键，并且需要指定pk2为自增主键。
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
	putPk.AddPrimaryKeyColumnWithAutoIncrement("pk2")
	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3")) //BINARY类型
	//2.PutRowChange
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = tableName
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("col1", "col1data1")
	putRowChange.AddColumn("col2", int64(100))
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

	//3.PutRowRequest
	putRowRequest := new(tablestore.PutRowRequest)
	putRowRequest.PutRowChange = putRowChange

	//4.PutRow
	_, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("put row failed with error:", err)
	} else {
		fmt.Println("put row finished")
	}
}

//--------------------------------------------------------------------------------------------------------------------







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

