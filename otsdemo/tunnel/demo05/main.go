package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
)

func main() {
	req := &tunnel.DeleteTunnelRequest {
		TableName: constants.DemoTable,
		TunnelName: constants.DemoTableTunnel2,
	}
	_, err := bootstrap.TunnelClient.DeleteTunnel(req)
	if err != nil {
		log.Fatal("delete test tunnel failed", err)
	}
}
