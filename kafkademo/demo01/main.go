package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)


//todo 演示一个生产者
func main() {


	//1.新建一个sarama配置实例
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll  //WaitForAll等待所有同步副本提交，然后再响应。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //返回一个分区器，它每次都选择一个随机分区。
	config.Producer.Return.Successes = true

	//2.新建一个同步生产者
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	
	//3.定义一个生产消息，包括Topic、消息内容、

	//3.1 发送消息01
	msg := &sarama.ProducerMessage{}
	msg.Topic = "topic-fast"
	msg.Key = sarama.StringEncoder("msg01")
	msg.Value = sarama.StringEncoder("hello world...")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)


    //3.2 发送消息02
	msg2 := &sarama.ProducerMessage{}
	msg2.Topic = "topic-fast"
	msg2.Key = sarama.StringEncoder("msg02")
	msg2.Value = sarama.StringEncoder("hello world2...")
	pid2, offset2, err := client.SendMessage(msg2)

	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	fmt.Printf("pid2:%v offset2:%v\n", pid2, offset2)
}