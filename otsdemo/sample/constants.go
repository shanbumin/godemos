package sample

//整体策划表
const  DemoTable="sbm_demos" //测试表名

//表结构操作相关的测试表
const TestTable ="sbm_test"
const Test2Table  = "sbm_test2"
const Test2TableDefinedcol1Index  =  Test2Table+"_gsi_definedcol1"
const Test3Table  ="sbm_table3" //主键自增测试
const Test4Table  ="sbm_table4"  //用来测试条件更新




//全局二级索引表
const GSI1Table  =  "sbm_call_record"
const GSI1CalledNumberIndex =GSI1Table+"_gsi_called_number"
const GSI1BaseStationNumber1Index =GSI1Table+"_gsi_base_station_number1"
const GSI1BaseStationNumber2Index =GSI1Table+"_gsi_base_station_number2"

const GSI2Table  =  "sbm_global_secondary_index"
const GSI2Definedcol1Index  = GSI2Table+"_gsi_definedcol1"


//sdk操作相关表
const BatchName  =  "sbm_batch"



const SingleName = "sbm_single"
const SearchIndexName  = "sbm_search_index"
const SearchIndex1 ="idx001"








