package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

//表名默认就是结构体名称的复数
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Age2         int64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
//将 User 的表名设置为 `profiles`
func (u User) TableName() string {
	return "profiles"
}
//---------------------------------------------------------


//parseTime是查询结果是否自动解析为时间
//loc是MySQL的时区设置
func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1)/gormdemo?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		fmt.Println(err)
	}
	defer db.Close()
	//建表一
	//db.CreateTable(&User{}) //默认是users
	//建表二
	db.CreateTable(&User{}) //设置TableName指明表名 profiles

	//建表一
	//db.Table("deleted_users").CreateTable(&User{})  // 使用User结构体创建名为`deleted_users`的表


}
