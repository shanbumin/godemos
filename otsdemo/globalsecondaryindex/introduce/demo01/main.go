package main


//查询号码234567的所有主叫话单。

//表格存储的模型是对所有行按照主键进行排序，并且提供顺序扫描（getRange）接口，
//所以只需要在调用getRange接口时，将cell_number列（主叫号码）的最大及最小值均设置为234567，
//start_time列（通话发生时间）的最小值设置为0，最大值设置为INT_MAX，对数据表进行扫描即可。

//todo 解决手法是 在数据表中进行主键的范围查询:多行数据操作-范围读（GetRange）
func main()  {




}
