package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"os"
	"os/signal"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"syscall"
	"time"
)

type Config struct {
	Endpoint  string
	Instance  string
	TableName string
	AkId      string
	AkSecret  string
}

var testConfig = Config{
	Endpoint:  "<Your instance endpoint>",
	Instance:  "<Your instance name>",
	TableName: "<Your table name>",
	AkId:      "<Your ak id>",
	AkSecret:  "<Your ak secret>",
}

func main() {
	//创建客户端
	//tunnelClient := tunnel.NewTunnelClient(testConfig.Endpoint, testConfig.Instance, testConfig.AkId, testConfig.AkSecret)
	tunnelClient:=bootstrap.TunnelClient
	//创建通道
	//tunnelName := "exampleTunnel"
	//req := &tunnel.CreateTunnelRequest{
	//	TableName:  testConfig.TableName,
	//	TunnelName: tunnelName,
	//	Type:       tunnel.TunnelTypeBaseStream,
	//}
	//resp, err := tunnelClient.CreateTunnel(req)
	//if err != nil {
	//	log.Fatal("create test tunnel failed", err)
	//}
	//log.Println("tunnel id is", resp.TunnelId)

	//start consume tunnel
	workConfig := &tunnel.TunnelWorkerConfig{
		HeartbeatTimeout: 1 * time.Second,//worker同Tunnel服务的心跳超时时间，通常使用默认值即可。
		HeartbeatInterval: 3* time.Second,//worker发送心跳的频率，通常使用默认值即可。
		ProcessorFactory: &tunnel.SimpleProcessFactory{
			CustomValue: "user defined interface{} value",
			ProcessFunc: exampleConsumeFunction,
			ShutdownFunc: func(ctx *tunnel.ChannelContext) {
				fmt.Println("shutdown hook")
			},
		},
	}

	daemon := tunnel.NewTunnelDaemon(tunnelClient,constants.DemoTableTunnelID, workConfig)
	go func() {
		err:= daemon.Run()
		if err != nil {
			log.Fatal("tunnel worker fatal error: ", err)
		}
	}()

	{
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-stop
		daemon.Close()
	}
}

func exampleConsumeFunction(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	fmt.Println("user-defined information", ctx.CustomValue)
	for _, rec := range records {
		time.Sleep(1 * time.Second)
		fmt.Println("tunnel record detail:", rec.String())
		fmt.Println("-----\r\n")
	}
	fmt.Println("a round of records consumption finished") //一轮消费结束
	return nil
}
