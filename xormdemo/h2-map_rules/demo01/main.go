package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"log"
	"time"
	"xorm.io/xorm"
)

//联合普通索引
//唯一索引
//普通索引
//默认值
//非空
//...


//todo 具体的 Tag 规则如下，另 Tag 中的关键字均不区分大小写，但字段名根据不同的数据库是区分大小写：
//@link https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-02/4.columns.html
//@link https://blog.csdn.net/wangzehui_1/article/details/79162444  tinyint(1)与tinyint(3)有区别吗
//@link  https://www.jianshu.com/p/b15cdc114458  decimal
type Student struct {
	Id        int64  //id主键自增
	Name     string  `xorm:"'user_name' varchar(255) notnull  index  default '' comment('姓名')"` //姓名，普通索引
	Email    string  `xorm:"varchar(200) notnull unique default '' comment('邮箱')"` //邮箱，唯一
	First    string  `xorm:"varchar(100) notnull default '' index('fs') comment('一号')"` //模拟演示联合普通索引字段一
	Second    string  `xorm:"varchar(100) notnull default '' index('fs') comment('二号')"` //模拟演示联合普通索引字段二
	Age       uint8     `xorm:"tinyint(3) notnull default 0 comment('年纪')"` //年纪
	Url       string  `xorm:"-"`
	Money     decimal.Decimal  `xorm:"decimal(10,2) comment('钱')"`
	Salary    float32
	Salary2   float64
	Status    uint8    `xorm:"tinyint(1) notnull default 0 comment('状态')"`
	File      []uint8  `xorm:"comment('文件流')"`
	Alived    bool      `xorm:"notnull default 0 comment('活着吗')"`
	//行为符特训
	Version   int64   `xorm:"version"` //todo 这个Field将会在insert时默认为1，每次更新自动加1
	SpecialCreatedAt   int `xorm:"created default 0"` //todo 这里的created不是字段名，而是行为符
	CreatedAt time.Time `xorm:"created"` //todo 注意不要用单引号括起来created,单引号是用来括字段名，当字段名与行为符冲突的时候额
	UpdatedAt time.Time  `xorm:"updated"` //todo 这里的updated是行为符而不是字段名
	DeletedAt time.Time `xorm:"deleted"` //todo 这里的deleted是行为符(此特性会激发软删除)
	Desc      Desc    `xorm:"json comment('描述')"`
}

func (m *Student) TableName() string {
	return "student"
}

type Desc struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
}

//---------------------------------------------------------------------------------------------------------------------
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

	file:=make([]uint8,10)
	file[0]=1
	file[1]=2
	//插入
	stu01:=&Student{
		Name:             "sam",
		Email:            "sam2@qq.com",
		First:            "步惊云",
		Second:           "毛阿敏",
		Age:              18,
		Url:              "http://www.baidu.com",
		Money:            decimal.NewFromFloat(100.89),
		Salary:           18.9,
		Salary2:          18.9,
		Status:           1,
		File:            file,
		Alived:           true,
		Desc:             Desc{"love","what is a love."},
	}
	_,err=engine.Insert(stu01)
	if err !=nil{
		fmt.Println(err)
	}

}

//todo 另外有如下几条自动映射的规则：
//1.如果field名称为Id而且类型为int64并且没有定义tag，则会被xorm视为主键，并且拥有自增属性。
  //如果想用Id以外的名字或非int64类型做为主键名，必须在对应的Tag上加上xorm:"pk"来定义主键，加上xorm:"autoincr"作为自增。
  //这里需要注意的是，有些数据库并不允许非主键的自增属性。

//2.string类型默认映射为varchar(255)，如果需要不同的定义，可以在tag中自定义，如：varchar(1024)

//...





