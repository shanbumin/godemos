package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"otsdemo/bootstrap"
	"otsdemo/constants"
)

//todo 其中ProcessorFactory为用户注册消费callback函数以及其他信息的接口，建议使用SDK中自带SimpleProcessorFactory实现。
//todo DefaultLogConfig是TunnelWorkerConfig和SimpleProcessFactory使用的默认日志配置。
//todo DefaultSyncer是TunnelWorkerConfig和SimpleProcessFactory使用的默认日志轮转配置
//@link https://help.aliyun.com/document_detail/102495.html?spm=a2c4g.11186623.6.663.5e5a7562iytFwC  配置TunnelWorkerConfig
func main() {


	//根据业务自定义数据消费Callback函数，开始自动化的数据消费。
	//配置callback到SimpleProcessFactory，配置消费端TunnelWorkerConfig。
	workConfig := &tunnel.TunnelWorkerConfig{
		HeartbeatTimeout:tunnel.DefaultHeartbeatTimeout,//worker同Tunnel服务的心跳超时时间，通常使用默认值即可。
		HeartbeatInterval:tunnel.DefaultHeartbeatInterval,//worker发送心跳的频率，通常使用默认值即可。
        //ChannelDialer:? //tunnel下消费连接建立接口，通常使用默认值即可。
		LogConfig:&tunnel.DefaultLogConfig,//zap日志配置，默认值为DefaultLogConfig。
		LogWriteSyncer:tunnel.DefaultSyncer,//zap日志轮转配置，默认值为DefaultSyncer。
		ProcessorFactory: &tunnel.SimpleProcessFactory{ //消费连接上具体处理器产生接口，通常使用callback函数初始化SimpleProcessFactory即可。
			CustomValue: "sam is a good man.",//用户自定义信息，会传递到ProcessFunc和ShutdownFunc中的ChannelContext参数中。
			CpInterval: tunnel.DefaultCheckpointInterval,//Worker记录checkpoint的间隔，CpInterval<=0时会使用DefaultCheckpointInterval。
			ProcessFunc: exampleConsumeFunction,
			//ShutdownFunc:? //worker退出时的同步调用callback。
			Logger:nil,//日志配置，Logger为nil时会使用DefaultLogConfig初始化logger。
		},
	}
	//使用TunnelDaemon持续消费指定tunnel。
	daemon := tunnel.NewTunnelDaemon(bootstrap.TunnelClient,constants.DemoTableTunnel2ID, workConfig)
	log.Fatal(daemon.Run())

}


//根据业务自定义数据消费callback函数。
func exampleConsumeFunction(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	//fmt.Println("user-defined information", ctx.CustomValue)
	for k, rec := range records {
		fmt.Println(k,"——", rec.String())
	}
	//fmt.Println("a round of records consumption finished")
	return nil
}