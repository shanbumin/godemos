package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

//教师详细信息
type Detail struct {
	Id        int64
	Email     string
	Addr      string
	Tel       string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

//学生
type Student struct {
	Id        int64
	Name      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"` //软删除字段
}

//教师
type Teacher struct {
	Id        int64
	Name      string
	DetailId  int64     `xorm:"index notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

//课程
type Course struct {
	Id        int64
	Name      string
	TeacherId int64     `xorm:"index not null"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

//成绩表
type Performance struct {
	Id        int64
	CourseId  int64 `xorm:"index notnull"`
	StudentId int64 `xorm:"index notnull"`
	Score     decimal.Decimal
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

type TeacherDetail struct {
	Teacher `xorm:"extends"`
	Detail  `xorm:"extends"`
}

//func (TeacherDetail) TableName() string {
//  //指定使用该结构体对象 进行数据库查询时，使用的表名
//  return "teacher"
//}

type CourseTeacher struct {
	Course  `xorm:"extends"`
	Teacher `xorm:"extends"`
}

func (CourseTeacher) TableName() string {
	return "course"
}

func main() {

	//一、获取引擎
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		logrus.Panicf("数据库连接错误，%v", err)
	}

	//设置日志显示
	engine.ShowSQL(true)
	engine.SetLogLevel(log.LOG_DEBUG)

	//设置连接池
	engine.SetMaxOpenConns(3)
	engine.SetMaxIdleConns(1)
	engine.SetConnMaxLifetime(12 * time.Hour)
	// 设置缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// engine.SetDefaultCacher(cacher)
	//测试连接
	engine.Ping()

	//二、将结构体同步到数据库
	//create database demo;
	//推荐生产环境按照传统的flyway的处理方式，
	//将每个版本的sql幂等处理、编号、归档，升级时各局点存储升级记录、做sql文件md5校验，避免混乱。
	//而通过sync2方式编写sql升级管理流程，每次升级需要编写升级代码，显然通过sql的配置文件的形式更方便。
	//err = engine.Sync2(new(Detail), new(Student), new(Teacher), new(Course), new(Performance))
	//if err != nil {
	//	logrus.Panicf("同步到数据库失败，%v", err)
	//}

	//三、插入基础数据
	/*

		tcher1 := &Teacher{Id: 1, Name: "卡卡西", DetailId: 1}
		tcher2 := &Teacher{Id: 2, Name: "凯", DetailId: 2}
		detail1 := &Detail{Id: 1, Tel: "卡卡卡卡卡", Addr: "卡卡卡", Email: "kakaka@sina.com"}
		detail2 := &Detail{Id: 2, Tel: "11111111", Addr: "木叶村", Email: "kai@sina.com"}


		stu1 := &Student{Id: 1, Name: "佐助"}
		stu2 := &Student{Id: 2, Name: "鸣人"}
		stu3 := &Student{Id: 3, Name: "小樱"}
		stu4 := &Student{Id: 4, Name: "雏田"}

		course1 := &Course{Name: "疾风手里剑", TeacherId: 1}
		course2 := &Course{Name: "基本体术", TeacherId: 2}

		perf1 := &Performance{CourseId: 1, StudentId: 1, Score: decimal.NewFromFloat(100)}
		perf2 := &Performance{CourseId: 1, StudentId: 2, Score: decimal.NewFromFloat(60)}
		perf3 := &Performance{CourseId: 1, StudentId: 3, Score: decimal.NewFromFloat(80)}

		engine.Insert(detail1, detail2, stu1, stu2, stu3, stu4, tcher1, tcher2, course1, course2, perf1, perf2, perf3)
	*/

	//四、根据主键获取一条记录的三种写法  todo 仔细查看sql输出，拼接的方式还是不一样的
	//theOne := &Student{Id: 1}
	//theOther := &Student{}
	//_, _ = engine.Get(theOne)
	//_, _ = engine.ID(1).Get(theOther)
	//_, _ = engine.Where("id=?", 1).Get(theOther)

	//五、软删除
	//_, err = engine.Delete(&Student{Name: "雏田"})
	//if err !=nil{
	//	fmt.Println(err)
	//}

	//六、单表，多条查询，值存入传入的对象中
	//6.1 返回数组
	fmt.Println("-----------------------------------------------------")
	stuArr1 := make([]Student, 0) //make声明slice
	engine.Find(&stuArr1)
	logrus.Infof("查询学生结构体数组：%+v\r\n", stuArr1)

	//6.2 返回map
	fmt.Println("-----------------------------------------------------")
	stuMap1 := make(map[int64]Student) //make声明Map
	engine.Find(&stuMap1)
	logrus.Infof("查询学生map：%v", stuMap1)

	//6.3 1对1查询（1教师：1教师详情）, 教师表  教师详情表  Cols("teacher.*", "detail.*")   Select("teacher.*,detail.*")
	fmt.Println("-----------------------------------------------------")
	//select t.id,t.name,d.email,d.addr,d.tel from teacher as t left join detail as d  on t.detail_id=d.id\G
	tcherDetails := make([]TeacherDetail, 0) //使用Table指定时，可不用编写结构体的TableName方法
	engine.Table("teacher").Alias("t").Select("t.id,t.name,d.email,d.addr,d.tel").Join("LEFT", []string{"detail", "d"}, "t.detail_id=d.id").Find(&tcherDetails)
	//engine.Table("teacher").Select("teacher.id,teacher.name,detail.email,detail.addr,detail.tel").Join("LEFT","detail", "teacher.detail_id=detail.id").Find(&tcherDetails)
	logrus.Infof("查询1对1（1教师：1教师详情）：%v", tcherDetails)

	//6.4   课程表中的teacher_id需要额外连表才能获知
	fmt.Println("-----------------------------------------------------")
	courseTchers := make([]CourseTeacher, 0) //声明数组
	engine.Join("LEFT", "teacher", "course.teacher_id=teacher.id").Find(&courseTchers)
	logrus.Infof("查询1对多（N课程：1老师）：%v\r\n", courseTchers)

	//6.5 多对多查询  todo 成绩表中的course_id和student_id需要我们额外连表才能获知
	fmt.Println("-----------------------------------------------------")
	performs := make([]Performance, 0)
	engine.Join("LEFT", "course", "performance.course_id=course.id").
		Join("left", "student", "performance.student_id=student.id").
		Find(&performs)
	logrus.Infof("查询多对多(N课程：N学生)：%v\r\n", performs)

	//6.6 不定义自己的集合，使用rows的封装类型返回
	fmt.Println("-----------------------------------------------------")
	newStu := new(Student)
	rows, err := engine.Rows(newStu)
		if err != nil {

			defer rows.Close()
			for rows.Next() {
				_ = rows.Scan(newStu)
				logrus.Infof("newStu:%v\r\n", newStu)
			}
		 }

	//6.7 执行sql1    [执行SQL查询]https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-08/index.html
	sql1 := "select * from student where id = ?"
	//queryRet, err := engine.Query(sql1, 1)
	queryRet, err := engine.QueryString(sql1, 1)
	if err != nil {
		fmt.Println(err)
	}
	logrus.Infof("使用sql查询结果：%v\r\n", queryRet)

	//6.8 执行sql2  [执行SQL命令] https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-09/index.html
	//sql2 := "insert into student (name,created_at,updated_at) values (?,now(),now())"
	//ret, err := engine.Exec(sql2, "日向宁次")
	//lastInsertId, _ := ret.LastInsertId()
	//effectedRows, _ := ret.RowsAffected()
	//logrus.Infof("执行sql命令结果,插入id：%v，影响行数：%v\r\n", lastInsertId, effectedRows)

	//七、事务 [事务处理]https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-10/index.html
	fmt.Println("-----------------------------------------------------")
	//txExample(engine)
	//八、事件  https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-12/index.html
	fmt.Println("-----------------------------------------------------")
	before := func(bean interface{}) {
		fmt.Println("before", bean)
	}
	after := func(bean interface{}) {
		fmt.Println("after", bean)
	}
	engine.Before(before).After(after).Get(new(Student))

}


func txExample(engine *xorm.Engine) {
	session := engine.NewSession()
	defer session.Close()

	//插入 教师详情
	detail := &Detail{
		Email: "zilaiye@sina.com",
		Addr:  "木叶村",
		Tel:   "999",
	}

	_, err := session.Insert(detail) //insert后会更新 detail变量中的相关字段，id、createdAt、updatedAt等
	if err != nil {
		session.Rollback() //插入失败回滚
		return
	}

	//插入 教师
	tcher := &Teacher{
		Name:     "自来也",
		DetailId: detail.Id,
	}
	_, err = session.Insert(tcher)
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		return
	}
}