package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/golang/protobuf/proto"
	"otsdemo/constants"
	"otsdemo/bootstrap"
)


//todo  IndexSchema=FieldSchemas+IndexSetting+IndexSort
//todo  FieldSchema=fieldName（必选）+fieldType（必选）+array（可选）+index（可选）+ indexOptions（可选） +analyzer（可选） + enableSortAndAgg（可选）+ store（可选）
//创建多元索引
func main() {
	tableName:= constants.DemoTable
	indexName:=tableName+"_index"
	client:= bootstrap.Client
	request := &tablestore.CreateSearchIndexRequest{}
	//1.数据表名称
	request.TableName = tableName
	//2.多元索引名称
	request.IndexName = indexName
	//3.索引的结构信息  IndexSchema
	indexSchema:=&tablestore.IndexSchema{}
	//-------------
	fieldName := &tablestore.FieldSchema{
		FieldName:        proto.String("name"),  // 设置字段名，使用proto.String用于获取字符串指针
		FieldType:        tablestore.FieldType_KEYWORD, // 设置字段类型  字符串
		Index:            proto.Bool(true),             // 设置开启索引
		EnableSortAndAgg: proto.Bool(true),             // 设置开启排序与统计功能
	}
	fieldAge := &tablestore.FieldSchema{
		FieldName:        proto.String("age"),
		FieldType:        tablestore.FieldType_LONG, //长整型
		Index:            proto.Bool(true),
		EnableSortAndAgg: proto.Bool(true),
	}
	fieldSalary := &tablestore.FieldSchema{
		FieldName:        proto.String("salary"),
		FieldType:        tablestore.FieldType_DOUBLE,
		Index:            proto.Bool(true),
		EnableSortAndAgg: proto.Bool(true),
	}
	fieldMarried := &tablestore.FieldSchema{
		FieldName:        proto.String("married"),
		FieldType:        tablestore.FieldType_BOOLEAN,
		Index:            proto.Bool(true),
		EnableSortAndAgg: proto.Bool(true),
	}
	fieldCreatedAt := &tablestore.FieldSchema{
		FieldName:        proto.String("created_at"),
		FieldType:        tablestore.FieldType_LONG,
		Index:            proto.Bool(true),
		EnableSortAndAgg: proto.Bool(true),
	}
	//--------
	schemas := []*tablestore.FieldSchema{}
	schemas = append(schemas, fieldName, fieldAge,fieldSalary,fieldMarried,fieldCreatedAt)
	indexSchema.FieldSchemas=schemas
	//indexSchema.IndexSetting=  //索引设置
	//indexSchema.IndexSort= //索引预排序设置(IndexSort)
	request.IndexSchema = indexSchema





	resp, err := client.CreateSearchIndex(request)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Println("CreateSearchIndex finished, requestId:", resp.ResponseInfo.RequestId)


	
}
