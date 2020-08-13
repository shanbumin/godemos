package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)
//https://github.com/etcd-io/etcd/blob/master/clientv3/client.go
//v3的客户端是放在了clientv3包下
func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		Username:  "",
		Password:  "",
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(client)

}

