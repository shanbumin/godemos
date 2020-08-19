package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

type Student struct {
	Id        int64 `xorm:"pk autoincr"`
	Age      int     `xorm:"notnull default 0 comment('年纪')"`
	Name     string  `xorm:"varchar(255) notnull unique  default '' comment('姓名')"`
	Url       string  `xorm:"index notnull index default '' comment('地址')"`
	Status    uint8    `xorm:"notnull default 0 comment('状态')"`
	Version   int64   `xorm:"version"`
	SpecialCreatedAt   int64 `xorm:"created"`
	CreatedAt time.Time `xorm:"created"` //todo 注意不要用单引号将行为符created括起来,单引号是用来括字段名的(当字段名与行为符冲突的时候)
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"` // 此特性会激发软删除
}

func (m *Student) TableName() string {
	return "student51"
}

func main() {

	driverName:="mysql"
	dataSourceName:="root:root@tcp(127.0.0.1:3306)/gwp?charset=utf8&parseTime=true"
	//第一步创建引擎
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err !=nil{
		log.Fatal(err)
	}
	//第二步 同步创建表
	err=engine.Sync2(new(Student))
	if err !=nil{
		log.Fatal(err)
	}



	//给Table设定一个别名
   stu:=&Student{}
   engine.Alias("s").Where("s.name = ?","sam").Get(stu)
   fmt.Println(stu)




}
