package main

import (
	"otsdemo/gsi/ready/install"
	"otsdemo/sample"
	"otsdemo/sdk/start"
	"time"
)

func main() {

	//先创建一个初始表
	install.CreateGSI1TableSample(start.Client,sample.GSI1Table) //场景演练需要
	install.CreateGSI2TableSample(start.Client,sample.GSI2Table) //sdk需要
	time.Sleep(3*time.Second)
	//插入初始化数据
	install.BatchWriteGSI1TableSample(start.Client,sample.GSI1Table)
	install.BatchWriteGSI2TableSample(start.Client,sample.GSI2Table)
}




