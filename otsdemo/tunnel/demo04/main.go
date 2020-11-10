package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
)

func main() {



	req := &tunnel.DescribeTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: constants.DemoTableTunnel2,
	}
	resp, err := bootstrap.TunnelClient.DescribeTunnel(req)
	if err != nil {
		log.Fatal("describe test tunnel failed", err)
	}
	log.Println("tunnel id is", resp.Tunnel.TunnelId)



}
