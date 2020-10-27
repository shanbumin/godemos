package sample
import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
)

//创建数据表
func CreateTableSample(client *tablestore.TableStoreClient, tableName string) {
	createTableRequest := new(tablestore.CreateTableRequest)

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
	createTableRequest.TableMeta = tableMeta


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
	createTableRequest.TableOption = tableOption

	//3.为数据表配置预留读吞吐量或预留写吞吐量。reservedThroughput
	//容量型实例中的数据表的预留读/写吞吐量只能设置为0，不允许预留。
	//默认值为0，即完全按量计费。
	//单位为CU。
	//当预留读吞吐量或预留写吞吐量大于0时，表格存储会根据配置为数据表预留相应资源，且数据表创建成功后，将会立即按照预留吞吐量开始计费，超出预留的部分进行按量计费。
	//当预留读吞吐量或预留写吞吐量设置为0时，表格存储不会为数据表预留相应资源。
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0 //预留读吞吐量
	reservedThroughput.Writecap = 0  //预留写吞吐量
	createTableRequest.ReservedThroughput = reservedThroughput

	//4.索引表的结构信息 IndexMeta
	//todo 略




	_, err := client.CreateTable(createTableRequest)
	if (err != nil) {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}

}


//创建数据表的同时创建全局二级索引
func CreateTableWithGlobalIndexSample(client *tablestore.TableStoreClient, tableName string,indexName string) {

	createTableRequest := new(tablestore.CreateTableRequest)

	//1.数据表的结构信息 TableMeta
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddDefinedColumn("definedcol1", tablestore.DefinedColumn_STRING)
	tableMeta.AddDefinedColumn("definedcol2", tablestore.DefinedColumn_INTEGER)
	createTableRequest.TableMeta = tableMeta

	//2.数据表的配置信息
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 1
	createTableRequest.TableOption = tableOption
	//3.为数据表配置预留读吞吐量或预留写吞吐量。reservedThroughput
	reservedThroughput := new(tablestore.ReservedThroughput)
	createTableRequest.ReservedThroughput = reservedThroughput
	//4.索引表的结构信息 IndexMeta
	indexMeta := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta.AddPrimaryKeyColumn("definedcol1") //设置数据表的definedcol1列作为索引表的主键。
	indexMeta.AddDefinedColumn("definedcol2") //设置数据表的definedcol2列作为索引表的属性列。
	indexMeta.IndexName = indexName
	createTableRequest.AddIndexMeta(indexMeta)




	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:",err)
	} else {
		fmt.Println("Create table finished")
	}




}


//创建表时，只需将自增的主键属性设置为AUTO_INCREMENT。
func CreateTableKeyAutoIncrementSample(client *tablestore.TableStoreClient,tableName string) {

	createTableRequest := new(tablestore.CreateTableRequest)
	//1.TableMeta
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumnOption("pk2", tablestore.PrimaryKeyType_INTEGER, tablestore.AUTO_INCREMENT)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	createTableRequest.TableMeta = tableMeta

	//2.TableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	createTableRequest.TableOption = tableOption

	//3.ReservedThroughput
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createTableRequest.ReservedThroughput = reservedThroughput




	client.CreateTable(createTableRequest)




}



//条件更新需要的测试表
func CreateTableConditionSample(client *tablestore.TableStoreClient,tableName string) {

	createTableRequest := new(tablestore.CreateTableRequest)
	//1.TableMeta
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("pk1", tablestore.PrimaryKeyType_STRING)
	tableMeta.AddPrimaryKeyColumn("pk2", tablestore.PrimaryKeyType_INTEGER)
	tableMeta.AddPrimaryKeyColumn("pk3", tablestore.PrimaryKeyType_BINARY)
	createTableRequest.TableMeta = tableMeta

	//2.TableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1
	tableOption.MaxVersion = 3
	createTableRequest.TableOption = tableOption

	//3.ReservedThroughput
	reservedThroughput := new(tablestore.ReservedThroughput)
	reservedThroughput.Readcap = 0
	reservedThroughput.Writecap = 0
	createTableRequest.ReservedThroughput = reservedThroughput




	client.CreateTable(createTableRequest)




}


func PutRowWithConditionSample(client *tablestore.TableStoreClient,tableName string) {

	putRowRequest := new(tablestore.PutRowRequest)

	//1.PutRowChange
	putRowChange := new(tablestore.PutRowChange)
	//1.1设置主键，必须按照创建数据表时的顺序添加主键，并且需要指定pk2为自增主键。
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
	putPk.AddPrimaryKeyColumn("pk2", int64(2))
	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3")) //BINARY类型
	putRowChange.PrimaryKey = putPk
	//1.2表名
	putRowChange.TableName = tableName
	//1.3 AddColumn
	putRowChange.AddColumn("col1", "col1data1")
	putRowChange.AddColumn("col2", int64(100))
	//1.4 SetCondition  todo sdk包存在bug，该值必须得传递，按理不传递应该默认设置成IGNORE
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

	putRowRequest.PutRowChange = putRowChange

	_, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("put row failed with error:", err)
	} else {
		fmt.Println("put row finished")
	}
}






//----------------------------------------------------------------------------------------------------------------------
//您可以使用UpdateTable接口来更新指定表的预留读/写吞吐量。
func UpdateTableSample(client *tablestore.TableStoreClient, tableName string) {

	updateTableReq := new(tablestore.UpdateTableRequest)

	//1.TableName
	updateTableReq.TableName = tableName
	//2.TableOption
	updateTableReq.TableOption = new(tablestore.TableOption)
	updateTableReq.TableOption.TimeToAlive = -1
	updateTableReq.TableOption.MaxVersion = 5


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
	describeTableReq := new(tablestore.DescribeTableRequest)
	//1.TableName

	describeTableReq.TableName = tableName



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
//Name：预定义列名称。 ColumnType：预定义列的数据类型。
//为数据表增加预定义列，预定义列分别为definedColumnName01（String类型）、definedColumnName02（INTEGER类型）、definedColumnName03（String类型）。
func AddDefinedColumn(client *tablestore.TableStoreClient, tableName string) {
	addDefinedColumnRequest := new(tablestore.AddDefinedColumnRequest)
	//1.数据表名称。
	addDefinedColumnRequest.TableName = tableName
	//2.为数据表预先定义一些非主键列以及其类型，可以作为索引表的属性列或索引列。
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName01",tablestore.DefinedColumn_STRING)
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName02",tablestore.DefinedColumn_INTEGER)
	addDefinedColumnRequest.AddDefinedColumn("definedColumnName03",tablestore.DefinedColumn_STRING)


	_, err := client.AddDefinedColumn(addDefinedColumnRequest)
	if (err != nil) {
		fmt.Println("Failed to Add DefinedColumn with error:", err)
	} else {
		fmt.Println("Add DefinedColumn finished")
	}
}


//删除数据表的预定义列definedColumnName01和definedColumnName02。
func DeleteDefinedColumn(client *tablestore.TableStoreClient, tableName string){
	deleteDefinedColumnRequest := new(tablestore.DeleteDefinedColumnRequest)
	//1.要删除的预定义列名称
	deleteDefinedColumnRequest.DefinedColumns = []string{"definedColumnName01","definedColumnName02"}
	//2.表名
	deleteDefinedColumnRequest.TableName = tableName

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

	deleteReq := new(tablestore.DeleteTableRequest)
	//1.表名
	deleteReq.TableName = tableName

	_, err := client.DeleteTable(deleteReq)
	if err != nil {
		fmt.Println("Failed to delete table with error:", err)
	} else {
		fmt.Println("Delete table finished")
	}
}

//--------------------------------------------------------------------------------------------------------------------




func PutRowWithKeyAutoIncrementSample(client *tablestore.TableStoreClient,tableName string) {

	putRowRequest := new(tablestore.PutRowRequest)

	//1.PutRowChange
	putRowChange := new(tablestore.PutRowChange)
	//1.1设置主键，必须按照创建数据表时的顺序添加主键，并且需要指定pk2为自增主键。
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pk1", "pk1value1")
	putPk.AddPrimaryKeyColumnWithAutoIncrement("pk2")
	putPk.AddPrimaryKeyColumn("pk3", []byte("pk3")) //BINARY类型
	putRowChange.PrimaryKey = putPk
	//1.2表名
	putRowChange.TableName = tableName
	//1.3 AddColumn
	putRowChange.AddColumn("col1", "col1data1")
	putRowChange.AddColumn("col2", int64(100))
	//1.4 SetCondition  todo sdk包存在bug，该值必须得传递，按理不传递应该默认设置成IGNORE
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

	putRowRequest.PutRowChange = putRowChange

	_, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("put row failed with error:", err)
	} else {
		fmt.Println("put row finished")
	}
}

//--------------------------------------------------------------------------------------------------------------------


//条件更新
//根据指定主键更新一行，如果指定的行存在，则更新成功，否则更新失败。
func ConditionRowUpdateSample(client *tablestore.TableStoreClient,tableName string){

	updateRowRequest := new(tablestore.UpdateRowRequest)

	//1.UpdateRowChange
	updateRowChange := new(tablestore.UpdateRowChange)
	//1.1 TableName
	updateRowChange.TableName = tableName
	//1.2 PrimaryKey
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	updateRowChange.PrimaryKey = updatePk
    //1.3 Column
	updateRowChange.DeleteColumn("col1")            //删除col1列。
	updateRowChange.PutColumn("col2", int64(77))    //新增col2列，值为77。
	updateRowChange.PutColumn("col4", "newcol3")    //新增col4列，值为"newcol3"。
	//1.4  期望指定行存在。
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)


	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := client.UpdateRow(updateRowRequest)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("条件更新成功")
	}


}

//根据指定主键删除一行，如果指定的行存在，且col2列的值为3，则更新成功，否则更新失败。
func ConditionColUpdateSample(client *tablestore.TableStoreClient,tableName string){
	deleteRowReq := new(tablestore.DeleteRowRequest)

	//1.DeleteRowChange
	deleteRowReq.DeleteRowChange = new(tablestore.DeleteRowChange)
	//1.1 TableName
	deleteRowReq.DeleteRowChange.TableName = tableName
	//1.2 PrimaryKey
	deletePk := new(tablestore.PrimaryKey)
	deletePk.AddPrimaryKeyColumn("pk1", "pk1value1")
	deletePk.AddPrimaryKeyColumn("pk2", int64(2))
	deletePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	deleteRowReq.DeleteRowChange.PrimaryKey = deletePk

	//1.3 期望行存在。
	deleteRowReq.DeleteRowChange.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
	//1.4 期望列col2的值为3。
	clCondition1 := tablestore.NewSingleColumnCondition("col2", tablestore.CT_EQUAL, int64(3))
	deleteRowReq.DeleteRowChange.SetColumnCondition(clCondition1)




	_, err := client.DeleteRow(deleteRowReq)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("条件更新成功")
	}
}



//写入数据时，使用updateRowChange接口对整型列做列值的增量变更，然后读取更新后的新值。
func UpdateRowWithIncrementColumn(client *tablestore.TableStoreClient, tableName string) {
	updateRowRequest := new(tablestore.UpdateRowRequest)

	//1.updateRowChange
	updateRowChange := new(tablestore.UpdateRowChange)
	//1.1 设置数据表名称。
	updateRowChange.TableName = tableName
	//1.2 PrimaryKey
	updatePk := new(tablestore.PrimaryKey)
	updatePk.AddPrimaryKeyColumn("pk1", "pk1increment")
	updatePk.AddPrimaryKeyColumn("pk2", int64(2))
	updatePk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	updateRowChange.PrimaryKey = updatePk

	//1.3 对列执行增量变更，例如+X，-X等。
	updateRowChange.IncrementColumn("col2", int64(30))
	//1.4 设置返回类型，返回进行原子计数操作的列的新值。
	updateRowChange.SetReturnIncrementValue()
	//1.5 SetCondition
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	//1.6 对于进行原子计数操作的列，设置需要返回列值的列名。
	updateRowChange.AppendIncrementColumnToReturn("col2")

	updateRowRequest.UpdateRowChange = updateRowChange

	resp, err := client.UpdateRow(updateRowRequest)
	if err != nil {
		fmt.Println("update failed with error:", err)
		return
	} else {
		fmt.Println("update row finished")
		fmt.Println(resp)
		fmt.Println(len(resp.Columns))
		fmt.Println(resp.Columns[0].ColumnName)
		fmt.Println(resp.Columns[0].Value)
		fmt.Println(resp.Columns[0].Timestamp)
	}
}

//-------------------------过滤器 ---------------------------

//构造SingleColumnValueFilter
func GetRowWithFilter(client *tablestore.TableStoreClient, tableName string) {
    //1.设置主键
	pk := new(tablestore.PrimaryKey)
	pk.AddPrimaryKeyColumn("pk1", "pk1value1")
	pk.AddPrimaryKeyColumn("pk2", int64(2))
	pk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	//2.设置过滤条件
	//ColumnName  过滤器中参考列的名称
	//ComparatorType  过滤器中的关系运算符
	//ColumnValue  过滤器中参考列的对比值。
	condition := tablestore.NewSingleColumnCondition("c1", tablestore.ComparatorType(tablestore.CT_EQUAL), "浙江")
	//当参考列在某行中不存在时，是否返回该行。类型为bool值，默认值为true，表示如果参考列在某行中不存在，则返回该行。
	//当设置FilterIfMissing为false时，如果参考列在某行中不存在，则不返回该行。
	condition.FilterIfMissing = true

	//3.组装查询条件
	criteria := &tablestore.SingleRowQueryCriteria{
		TableName:     tableName,
		PrimaryKey:    pk,
		MaxVersion:    1,
		Filter:        condition,
	}

	//4.查询
	getResp, err := client.GetRow(&tablestore.GetRowRequest{SingleRowQueryCriteria: criteria})
	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		colMap := getResp.GetColumnMap()
		fmt.Println("length is ", len(colMap.Columns))
		fmt.Println("get row col0 result is ", getResp.Columns[0].ColumnName, getResp.Columns[0].Value)
	}
}


//构造CompositeColumnValueFilter
func GetRowWithCompositeColumnValueFilter(client *tablestore.TableStoreClient, tableName string) {
	//1.设置主键
	pk := new(tablestore.PrimaryKey)
	pk.AddPrimaryKeyColumn("pk1", "pk1value1")
	pk.AddPrimaryKeyColumn("pk2", int64(2))
	pk.AddPrimaryKeyColumn("pk3", []byte("pk3"))
	//2.设置过滤条件
	//LogicOperator: 过滤器中的逻辑运算符
	filter := tablestore.NewCompositeColumnCondition(tablestore.LO_AND)
	filter1 := tablestore.NewSingleColumnCondition("c1", tablestore.CT_EQUAL, "浙江")
	filter2 := tablestore.NewSingleColumnCondition("c2", tablestore.CT_EQUAL, "杭州")
	filter.AddFilter(filter2)
	filter.AddFilter(filter1)

	//3.组装查询条件
	criteria := &tablestore.SingleRowQueryCriteria{
		TableName:  tableName,
		PrimaryKey: pk,
		MaxVersion: 1,
		Filter:     filter,
	}

	getResp, err := client.GetRow(&tablestore.GetRowRequest{SingleRowQueryCriteria: criteria})
	if err != nil {
		fmt.Println("getrow failed with error:", err)
	} else {
		colMap := getResp.GetColumnMap()
		fmt.Println("length is ", len(colMap.Columns))
		fmt.Println("get row col0 result is ", getResp.Columns[0].ColumnName, getResp.Columns[0].Value)
	}
}


