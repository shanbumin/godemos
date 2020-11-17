package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/tunnel/common"
)

//利用通道可以进行故障或者规模转移额
func main() {
	//展示分区通道列表
	req := &tunnel.DescribeTunnelRequest{
		TableName: constants.DemoTable,
		//TunnelName: constants.DemoTableTunnel,
		TunnelName: "earliest-2020-11-16_16-29-17",
	}

	tunnelId, channels, err := common.DescribeTunnel(req,bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(channels)
	//todo 暂停所有开启的通道
	err = common.SuspendAllOpenChannel(tunnelId, channels,bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}

}


