package main

import (
	"otsdemo/sample"
	"otsdemo/start"
)


//前提条件
//todo 已删除数据表上的索引表和多元索引。

func main() {

	//删除数据表
	sample.DeleteTableSample(start.Client,sample.TableName)




}
