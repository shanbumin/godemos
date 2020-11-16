package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/moka-mrp/sword-core/samutils"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"strconv"
	"time"
)



func  RandBool()bool{
	res:=samutils.MtRand(0,2)
	if res==0{
		return false
	}
	return true
}


var appids =map[int]string{
	1:"wx32d84d8f87a72062",
	2:"wxe81d9b20d97523c6",
	3:"wxe21d9b20d97523c1",
}

//
//var stringChinese    = []rune("白黑红蓝紫灰粉绿茶叶榨干茶色头发有点意思就是说你们都不知道怎么办了大半天的时间是什么时候回来的时候就可以了我在家的时候我就不知道了")
//var stringLen = len(stringChinese)
//
//func RandChinese(n int64) string {
//	rand.Seed(time.Now().UnixNano())
//	b := make([]rune, n)
//	for i := range stringChinese {
//		b[i] = stringChinese[rand.Intn(stringLen)]
//	}
//	return string(b)
//}

func main() {

//每次插入三个，一共插入5次



	for  i:=1;i<=3;i++{


		batchWriteReq := &tablestore.BatchWriteRowRequest{}
		//批量接口一次只能插入200行
		for j := 1; j <= 10; j++ {
			putRowChange := new(tablestore.PutRowChange)
			//1.TableName
			putRowChange.TableName = constants.DemoTable
			//2.主键
			fmt.Println(i*j)
			name:="sam"+strconv.Itoa(i*j)
			putPk := new(tablestore.PrimaryKey)
			putPk.AddPrimaryKeyColumn("appid",appids[i])
			putPk.AddPrimaryKeyColumn("openid",samutils.Md5(name))
			putRowChange.PrimaryKey = putPk
			//3.属性列
			putRowChange.AddColumn("name", name) //name唯一
			putRowChange.AddColumn("age",int64(j))
			putRowChange.AddColumn("salary",float64(j*100))
			putRowChange.AddColumn("married",RandBool())
			putRowChange.AddColumn("desc",samutils.RandStringWordL(5))
			putRowChange.AddColumn("img",[]byte(samutils.RandStringWordL(5)))
			putRowChange.AddColumn("created_at",int64(time.Now().Unix()))
			putRowChange.AddColumn("updated_at",int64(time.Now().Unix()))
			putRowChange.AddColumn("tags","[\"红\",\"黑\"]")
			putRowChange.AddColumn("nests","[{\"tag\":\"红\",\"score\":100.20},{\"tag\":\"黑\",\"score\":60}]")
			//4.条件更新
			putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

			batchWriteReq.AddRowChange(putRowChange)
		}
		response, err := bootstrap.Client.BatchWriteRow(batchWriteReq)
		if err != nil {
			fmt.Println("batch request failed with:", response)
		} else {
			fmt.Println("batch write row finished",i)
		}

		time.Sleep(5 * time.Second) //休息会再插入下一次



	}




}
