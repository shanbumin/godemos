package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel/protocol"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"time"
)


//创建一个从最开始的记录到2020-11-16 16:29:17期间的增量通道
func main() {

	//创建通道
	to := time.Date(2020, time.November,16,16,29,17,0, time.Local)

	req := &tunnel.CreateTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: "earliest-2020-11-16_16-29-17",
		Type:       tunnel.TunnelTypeStream, //创建增量类型
		StreamTunnelConfig: &tunnel.StreamTunnelConfig{
			Flag:      protocol.StartOffsetFlag_EARLIEST,
			EndOffset: uint64(to.UnixNano()),
		},
	}
	resp, err := bootstrap.TunnelClient.CreateTunnel(req)
	if err != nil {
		log.Fatal("create test tunnel failed", err)
	}
	log.Println("tunnel id is", resp.TunnelId)


}



