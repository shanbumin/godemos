package install

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

//private static final String INDEX0_NAME = "IndexOnBeCalledNumber";
//private static final String INDEX1_NAME = "IndexOnBaseStation1";
//private static final String INDEX2_NAME = "IndexOnBaseStation2";
//private static final String PRIMARY_KEY_NAME_1 = "CellNumber";
//private static final String PRIMARY_KEY_NAME_2 = "StartTime";
//private static final String DEFINED_COL_NAME_1 = "CalledNumber";
//private static final String DEFINED_COL_NAME_2 = "Duration";
//private static final String DEFINED_COL_NAME_3 = "BaseStationNumber";

//todo 一个号码不会同时与多个号码接通的，所以cell_number与start_time的组合是唯一的额
func  CreateGSI1TableSample(client *tablestore.TableStoreClient, tableName string){
	createTableRequest := new(tablestore.CreateTableRequest)
	//1.TableMeta
	tableMeta := new(tablestore.TableMeta)
	tableMeta.TableName = tableName
	tableMeta.AddPrimaryKeyColumn("cell_number", tablestore.PrimaryKeyType_INTEGER) //主叫号码
	tableMeta.AddPrimaryKeyColumn("start_time", tablestore.PrimaryKeyType_INTEGER) //通话发生时间
	//添加预定义列
	tableMeta.AddDefinedColumn("called_number", tablestore.DefinedColumn_INTEGER) //被叫号码
	tableMeta.AddDefinedColumn("duration", tablestore.DefinedColumn_INTEGER) //通话时长
	tableMeta.AddDefinedColumn("base_station_number", tablestore.DefinedColumn_INTEGER) //基站号码
	createTableRequest.TableMeta = tableMeta
	//2.IndexMeta
	//todo 说明 系统会自动进行索引列补齐。即把数据表的主键添加到索引列后，共同作为索引表的主键，所以索引表中有三列主键。
	indexMeta0 := new(tablestore.IndexMeta) //新建索引表Meta。
	indexMeta0.AddPrimaryKeyColumn("called_number")
	indexMeta0.IndexName = sample.GSI1CalledNumberIndex
	createTableRequest.AddIndexMeta(indexMeta0)

	indexMeta1:=new(tablestore.IndexMeta)
	indexMeta1.AddPrimaryKeyColumn("base_station_number")
	indexMeta1.AddPrimaryKeyColumn("start_time")
	indexMeta1.IndexName = sample.GSI1BaseStationNumber1Index
	createTableRequest.AddIndexMeta(indexMeta1)

	indexMeta2:=new(tablestore.IndexMeta)
	indexMeta2.AddPrimaryKeyColumn("base_station_number")
	indexMeta2.AddPrimaryKeyColumn("start_time")
	indexMeta2.AddDefinedColumn("duration")
	indexMeta2.IndexName = sample.GSI1BaseStationNumber2Index
	createTableRequest.AddIndexMeta(indexMeta2)


	//3.TableOption
	tableOption := new(tablestore.TableOption)
	tableOption.TimeToAlive = -1 //数据的过期时间，单位为秒，-1表示永不过期。带索引表的数据表数据生命周期必须设置为-1。
	tableOption.MaxVersion = 1 //保存的最大版本数，1表示每列上最多保存一个版本即保存最新的版本。带索引表的数据表最大版本数必须设置为1。
	createTableRequest.TableOption = tableOption

	//4.ReservedThroughput(保留吞吐量)
	reservedThroughput := new(tablestore.ReservedThroughput)
	createTableRequest.ReservedThroughput = reservedThroughput



	//5.create
	_, err := client.CreateTable(createTableRequest)
	if err != nil {
		fmt.Println("Failed to create table with error:", err)
	} else {
		fmt.Println("Create table finished")
	}
}