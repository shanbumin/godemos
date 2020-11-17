package common

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
)

/**
 * 返回通道分区列表
 */
func DescribeTunnel(req *tunnel.DescribeTunnelRequest, client tunnel.TunnelMetaApi) (string, []*tunnel.ChannelInfo, error) {
	resp, err := client.DescribeTunnel(req)
	if err != nil {
		return "", nil, err
	}
	//fmt.Println("该通道创建时间:",resp.Tunnel.CreateTime.Format("2006-01-02 15:04:05"))
	//遍历通道分区列表
	//for _, ch := range resp.Channels {
	//	fmt.Println("ChannelDetail:", ch)
	//}
	return resp.Tunnel.TunnelId, resp.Channels, nil
}



//暂停所有开启的通道分区列表
//@author sam@2020-11-17 11:38:36
func SuspendAllOpenChannel(tunnelId string, channels []*tunnel.ChannelInfo, client tunnel.TunnelMetaApi) error {
	channelsToSuspend := make([]*tunnel.ScheduleChannel, 0)
	for _, ch := range channels {
		if ch.ChannelStatus == "OPEN" {
			channelsToSuspend = append(channelsToSuspend, tunnel.SuspendChannel(ch.ChannelId))
		}
	}

	if  len(channelsToSuspend )>0{
		_, err := client.Schedule(&tunnel.ScheduleRequest{
			TunnelId: tunnelId,
			Channels: channelsToSuspend,
		})
		if err !=nil{
			return err
		}
	}

	return nil
}
//将所有处于等待状态的分区全部打开
func ScheduleAllWaitingChannel(tunnelId string, channels []*tunnel.ChannelInfo, client tunnel.TunnelMetaApi) error {
	channelsToOpen := make([]*tunnel.ScheduleChannel, 0)
	for _, ch := range channels {
		if ch.ChannelStatus == "WAIT" {
			channelsToOpen = append(channelsToOpen, tunnel.OpenChannel(ch.ChannelId))
		}
	}
	if  len(channelsToOpen )<=0{
		return nil
	}
	_, err := client.Schedule(&tunnel.ScheduleRequest{
		TunnelId: tunnelId,
		Channels: channelsToOpen,
	})
	return err
}

//恢复所有已经关闭的分区,可以让该分区中的消息重新再消费额
func ResumeAllSuspendedChannel(tunnelId string, channels []*tunnel.ChannelInfo, client tunnel.TunnelMetaApi) error {
	channelsToResume := make([]*tunnel.ScheduleChannel, 0)
	for _, ch := range channels {
		if ch.ChannelStatus == "CLOSING" {
			channelsToResume = append(channelsToResume, tunnel.ResumeChannel(ch.ChannelId))
		}
	}
	if len(channelsToResume)<=0{
	return nil
	}
	_, err := client.Schedule(&tunnel.ScheduleRequest{
		TunnelId: tunnelId,
		Channels: channelsToResume,
	})
	return err
}





