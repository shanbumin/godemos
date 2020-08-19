package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)




func main() {

	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8mb4&parseTime=true"
	//创建引擎
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
		log.Fatal(err)
	}

	//可以获取到数据库中所有的表，字段，索引的信息。
	tables,_:=engine.DBMetas()

	for _,v:=range tables{
        fmt.Printf("%+v\r\n",v.Name)
	}

}
