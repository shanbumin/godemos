package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

//类似于TermQuery，但是TermsQuery可以指定多个查询关键词，查询匹配这些词的数据。多个查询关键词中只要有一个词精确匹配，该行数据就会被返回，等价于SQL中的In。
func main() {
	//1.初始化对接
	client:=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
	//2.多词精确查询
	sample.TermsQuery(client,"sbm_search_index","idx001")
}
