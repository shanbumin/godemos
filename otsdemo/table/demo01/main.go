package main

import (
	"otsdemo/sample"
	"otsdemo/start"
)


//创建数据表
//todo  使用CreateTable接口创建数据表时，需要指定数据表的结构信息和配置信息，高性能实例中的数据表还可以根据需要设置预留读/写吞吐量。
//      创建数据表的同时支持创建一个或者多个索引表。

//todo  说明
//      创建表后需要几秒钟进行加载，在此期间对该表的读/写数据操作均会失败。应用程序应该等待表加载完毕后再进行数据操作。
//      创建表时[必须]指定表的主键。主键包含1~4个主键列，每一个主键列都有名称和类型。


func main() {
    //todo 创建数据表（不带索引）
    //todo 创建一个含有2个主键列，预留读/写吞吐量为(0, 0)的数据表。
    sample.CreateTableSample(start.Client,sample.TableName)
}


