package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel/protocol"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"time"
)



//从最近到最近到未来某个时间段
func main() {

	//创建通道
	to := time.Date(2088, time.November,16,16,29,17,0, time.Local)

	req := &tunnel.CreateTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: "latest-2088-11-16_16-29-17",
		Type:       tunnel.TunnelTypeStream, //创建增量类型
		StreamTunnelConfig: &tunnel.StreamTunnelConfig{
			Flag:      protocol.StartOffsetFlag_LATEST,
			EndOffset:   uint64(to.UnixNano()),
		},
	}
	resp, err := bootstrap.TunnelClient.CreateTunnel(req)
	if err != nil {
		log.Fatal("create test tunnel failed", err)
	}
	log.Println("tunnel id is", resp.TunnelId)


}



