package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/tunnel/common"
)

func main() {

	req := &tunnel.DescribeTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: constants.DemoTableTunnel,
	}

	tunnelId, channels, err := common.DescribeTunnel(req,bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}


	//todo 打开所有等待状态的通道分区
	err = common.ScheduleAllWaitingChannel(tunnelId, channels,bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}

}


