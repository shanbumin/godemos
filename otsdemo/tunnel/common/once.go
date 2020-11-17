package common

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
)
type UserCheckpointer interface {
	//指定主键和列名，返回该行之前的  通道分区ID(channelId), 序列信息(sequenceInfo), 属性列(map)
	getCheckpoint(id string, colNameToGet []string) (channelId string, sequenceInfo *tunnel.SequenceInfo, valueMap map[string]interface{}, err error)
	//更新指定列的channelId，sequenceInfo，condition是乐观锁的条件期望，valueMap是属性列map
	updateCheckpoint(id, channelId string, sequenceInfo *tunnel.SequenceInfo, condition *tablestore.RowCondition, valueMap map[string]interface{}) error
}

//消费函数
func ExactlyOnceIngestionFinalState(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	checkpointer := ctx.CustomValue.(UserCheckpointer)

	for _, rec := range records {
		//1.过滤掉base data record
		if rec.SequenceInfo == nil { //增量数据才有SequenceInfo
			continue
		}

		//2.获取分区主键的值
		id := rec.PrimaryKey.PrimaryKeys[0].Value.(string)
		//3.指定主键和列名，返回该行之前的  通道分区ID(channelId), 序列信息(sequenceInfo), 属性列(map)
		_, instateSeq, valueMap, err := checkpointer.getCheckpoint(id, nil)
		if err != nil {
			return err
		}
		//4.检查即将要消费的记录的最终状态
		duplicated, condition := CheckRecordFinalState(rec.SequenceInfo, instateSeq)
		if duplicated {
			continue //skip
		}
		//todo do something with valueMap
		//fmt.Println("map size", len(valueMap))
		err = checkpointer.updateCheckpoint(id, "", rec.SequenceInfo, condition, valueMap)
		if err != nil {
			return err
		}
	}
	fmt.Println("a round of records consumption finished")
	return nil
}

//即将插入的Seq   旧的Seq
//返回 是否重复的
func CheckRecordFinalState(incomingSeq, instateSeq *tunnel.SequenceInfo) (duplicated bool, condition *tablestore.RowCondition) {
	condition = new(tablestore.RowCondition)
	if instateSeq == nil { //数据行不存在
		condition.RowExistenceExpectation = tablestore.RowExistenceExpectation_EXPECT_NOT_EXIST
		return
	}
	if !tunnel.StreamRecordSequenceLess(instateSeq, incomingSeq) {
		duplicated = true
		return
	}
	condition.RowExistenceExpectation = tablestore.RowExistenceExpectation_EXPECT_EXIST
	seqBuf, _ := json.Marshal(instateSeq)
	condition.ColumnCondition = tablestore.NewSingleColumnCondition("SequenceInfo", tablestore.CT_EQUAL, seqBuf)
	return
}


