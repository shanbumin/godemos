package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

/*
批量读接口用于一次请求读取多行数据，也支持一次对多个数据表进行读取。
BatchGetRow由多个GetRow子操作组成。构造子操作的过程与使用GetRow接口时相同，也支持使用过滤器。

批量读取的所有行采用相同的参数条件，例如ColumnsToGet=[colA]，则要读取的所有行都只读取colA列。

BatchGetRow的各个子操作独立执行，表格存储会分别返回各个子操作的执行结果。

由于批量读取可能存在部分行失败的情况，失败行的错误信息在返回的BatchGetRowResponse中，但并不抛出异常。
因此调用BatchGetRow接口时，需要检查返回值，判断每行的状态是否成功。

 */



func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.批量读
	sample.BatchGetRowSample(client,"sbm_batch")
	
}
