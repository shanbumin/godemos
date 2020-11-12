package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
)



//todo: 消费延迟监控
//通道服务通过DescribeTunnel API提供了客户端消费数据延迟（即当前消费到的数据的时间点）信息，并在控制台提供了通道数据消费监控。

//StreamId:"sbm_demos_1605062436695529"
func main() {



	req := &tunnel.DescribeTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: constants.DemoTableTunnel,
	}
	resp, err := bootstrap.TunnelClient.DescribeTunnel(req)
	if err != nil {
		log.Fatal("describe test tunnel failed", err)
	}
	//log.Println("tunnel id is", resp.Tunnel.TunnelId)

	fmt.Printf("%#v\r\n",resp.Tunnel)

}
