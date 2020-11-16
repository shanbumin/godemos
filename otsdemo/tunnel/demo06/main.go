package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"time"
)

//只能针对某个区间创建增量通道===>比如创建一个16号的增量通道，注意之后到了17号，虽然是增量通道，这个时候也不会往这个通道同步的额
func main() {

	//创建通道
	from := time.Date(2020, time.November, 16,0, 0, 0, 0, time.Local)
	to := time.Date(2020, time.November,16,23,59,59,59, time.Local)

	req := &tunnel.CreateTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: "2020-11-16_0-0-0-2020-11-16_23-59-59",
		Type:       tunnel.TunnelTypeStream, //创建增量类型
		StreamTunnelConfig: &tunnel.StreamTunnelConfig{
			StartOffset: uint64(from.UnixNano()),
			EndOffset:   uint64(to.UnixNano()),
		},
	}
	bootstrap.TunnelClient.CreateTunnel(req)

	//创建第二个

	to = time.Date(2020, time.November,16,16,29,17,0, time.Local)

	req = &tunnel.CreateTunnelRequest{
		TableName: constants.DemoTable,
		TunnelName: "2020-11-16_0-0-0-2020-11-16_16-29-17",
		Type:       tunnel.TunnelTypeStream, //创建增量类型
		StreamTunnelConfig: &tunnel.StreamTunnelConfig{
			StartOffset: uint64(from.UnixNano()),
			EndOffset:   uint64(to.UnixNano()),
		},
	}
	bootstrap.TunnelClient.CreateTunnel(req)





}



