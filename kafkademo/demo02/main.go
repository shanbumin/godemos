package main

import (
	"sync"
	"github.com/Shopify/sarama"
	"fmt"
)

var wg sync.WaitGroup //等待所有的partitions消费完毕再结束

func main() {
	//1.创建consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}
	fmt.Println("connnect success...")
	defer consumer.Close()

	//2.选取topic
	partitions, err := consumer.Partitions("topic-fast")
	if err != nil {
		fmt.Println("geet partitions failed, err:", err)
		return
	}

	//3.消费吧
	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("topic-fast", p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func(){
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}


