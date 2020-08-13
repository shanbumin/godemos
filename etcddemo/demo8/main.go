package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)


//Revision问题

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		putOp clientv3.Op
		getOp clientv3.Op
		opResp clientv3.OpResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	// kv.Do(op)
	// kv.Put
	// kv.Get
	// kv.Delete

	// 创建OpPut: operation
	putOp = clientv3.OpPut("/cron/jobs/kill", "shit")
	// 执行OpPut
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入Revision:", opResp.Put().Header.Revision)
	// 创建OpGet
	getOp = clientv3.OpGet("/cron/jobs/kill")
	// 执行OP
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	// 打印
	fmt.Println("CreateRevision",opResp.Get().Kvs[0].CreateRevision,"ModRevision:", opResp.Get().Kvs[0].ModRevision)	// create rev == mod rev
	fmt.Println("Value:", string(opResp.Get().Kvs[0].Value))
}
