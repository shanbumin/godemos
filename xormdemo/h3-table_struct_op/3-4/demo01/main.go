package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)





type Student struct {
	Id        int64 `xorm:"pk autoincr"`
	Age      int     `xorm:"notnull default 0 comment('年纪')"`
	Name     string  `xorm:"varchar(255) notnull unique  default '' comment('姓名')"`
	Sex       bool    `xorm:"default 0 comment('性别')"`
	Url       string  `xorm:"index notnull index default '' comment('地址')"`
	Status    uint8    `xorm:"notnull default 0 comment('状态')"`
	Version   int64   `xorm:"version"`
	CreatedAt time.Time `xorm:"'created'"`
	UpdatedAt time.Time  `xorm:"'updated'"`
	DeletedAt time.Time `xorm:"'deleted'"` // 此特性会激发软删除
}

func (m *Student) TableName() string {
	return "student34"
}

//@link https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-03/4.sync.html
//todo  当表中有数据之后，增加新字段该如何处理呢
//todo  insert into student34  (name,age,status) values("sam",18,1);

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

}
