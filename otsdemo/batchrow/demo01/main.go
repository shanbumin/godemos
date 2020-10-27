package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)


//todo  准备
//sbm_batch
//pk1   string
//pk2   integer
//pk3   binary


/*
批量写接口用于在一次请求中进行批量的写入操作，也支持一次对多个数据表进行写入。
BatchWriteRow操作由多个PutRow、UpdateRow、DeleteRow子操作组成，构造子操作的过程与使用PutRow接口、UpdateRow接口和DeleteRow接口时相同，也支持使用条件更新。

BatchWriteRow的各个子操作独立执行，表格存储会分别返回各个子操作的执行结果。

由于批量写入可能存在部分行失败的情况，失败行的Index及错误信息在返回的BatchWriteRowResponse中，但并不抛出异常。因此调用BatchWriteRow接口时，需要检查返回值，判断每行的状态是否成功；
如果不检查返回值，则可能会忽略掉部分操作的失败。

当服务端检查到某些操作出现参数错误时，BatchWriteRow接口可能会抛出参数错误的异常，此时该请求中所有的操作都未执行。

 */

func main() {


	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.批量写
	sample.BatchWriteRowSample(client,"sbm_batch")
	
}
