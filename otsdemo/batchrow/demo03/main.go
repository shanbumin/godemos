package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

/*

范围读接口用于读取一个主键范围内的数据。

范围读接口支持按照确定范围进行正序读取和逆序读取，可以设置要读取的行数。如果范围较大，已扫描的行数或者数据量超过一定限制，会停止扫描，并返回已获取的行和下一个主键信息。
您可以根据返回的下一个主键信息，继续发起请求，获取范围内剩余的行。

GetRange操作可能在如下情况停止执行并返回数据。
扫描的行数据大小之和达到4 MB。
扫描的行数等于5000。
返回的行数等于最大返回行数。
当前剩余的预留读吞吐量已全部使用，余量不足以读取下一条数据。
说明 表格存储表中的行默认是按照主键排序的，而主键是由全部主键列按照顺序组成的，所以不能理解为表格存储会按照某列主键排序，这是常见的误区。

*/



func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.范围读
	sample.GetRangeSample(client,"sbm_batch")
	
}
