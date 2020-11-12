package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
)

//todo  Tunnel
//todo  TunnelWorkerConfig  = HeartbeatInterval+...
//todo  TunnelWorker(客户端)

//OTSTunnelExpired


//CreateTunnel操作为某张数据表创建一个通道，一张数据表上可以创建多个通道。在创建通道时需要指定数据表名称、通道名称和通道类型。
func main() {

	//创建通道
	req := &tunnel.CreateTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: constants.DemoTableTunnel,
		Type:       tunnel.TunnelTypeBaseStream, //创建全量加增量类型的Tunnel。
	}
	resp, err := bootstrap.TunnelClient.CreateTunnel(req)
	if err != nil {
		log.Fatal("create test tunnel failed", err)
	}
	log.Println("tunnel id is", resp.TunnelId)

}



