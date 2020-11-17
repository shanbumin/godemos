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
		TunnelName: "earliest-2020-11-16_16-29-17",
	}

	tunnelId, channels, err := common.DescribeTunnel(req,bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}

	//todo 恢复所有已经关闭的分区
	err = common.ResumeAllSuspendedChannel(tunnelId, channels, bootstrap.TunnelClient)
	if err != nil {
		log.Fatal(err)
	}

}
