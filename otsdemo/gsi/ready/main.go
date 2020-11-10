package main

import (
	"otsdemo/gsi/ready/install"
	"otsdemo/constants"
	"otsdemo/bootstrap"
	"time"
)

func main() {

	//先创建一个初始表
	install.CreateGSI1TableSample(bootstrap.Client, constants.GSI1Table) //场景演练需要
	install.CreateGSI2TableSample(bootstrap.Client, constants.GSI2Table) //sdk需要
	time.Sleep(3*time.Second)
	//插入初始化数据
	install.BatchWriteGSI1TableSample(bootstrap.Client, constants.GSI1Table)
	install.BatchWriteGSI2TableSample(bootstrap.Client, constants.GSI2Table)
}




