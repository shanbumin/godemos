package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"otsdemo/sample"
)

//初始化``TableStoreClient``实例。
//endPoint是表格存储服务的地址（例如'https://instance.cn-hangzhou.ots.aliyun.com:80'），必须以'https://'开头。
//accessKeyId是访问表格存储服务的AccessKeyID，通过官方网站申请或通过管理员获取。
//accessKeySecret是访问表格存储服务的AccessKeySecret，通过官方网站申请或通过管理员获取。
//instanceName是要访问的实例名，通过官方网站控制台创建或通过管理员获取。
//func NewClient(endPoint, instanceName, accessKeyId, accessKeySecret string, options ...ClientOption) *TableStoreClient

func main() {
//1.初始化对接
client:=tablestore.NewClient("https://tulong.cn-shanghai.ots.aliyuncs.com",  "tulong", "LTAI4G4nT4uRrmdmpz2XGkpV", "eBRDHMUr1NhuBBlqno1U7Hsi5adZ5O")


//2.创建表

//说明：根据指定的表结构信息创建数据表。
//request是CreateTableRequest类的实例，它包含TableMeta和TableOption以及ReservedThroughput。
//请参见TableMeta类的文档。
//当创建一个数据表后，通常需要等待几秒钟时间使partition load完成，才能进行各种操作。
//返回：CreateTableResponse
//CreateTable(request *CreateTableRequest) (*CreateTableResponse, error)


   //todo 创建数据表（不带索引）
   //todo 创建一个含有2个主键列，预留读/写吞吐量为(0, 0)的数据表。
  sample.CreateTableSample001(client,"t1")



}