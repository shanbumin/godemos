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





func main() {
	//创建客户端
	tunnelClient:=bootstrap.TunnelClient
	var daemon *tunnel.TunnelWorkerDaemon
	//start consume tunnel
	workConfig := &tunnel.TunnelWorkerConfig{
		HeartbeatTimeout: 1 * time.Second,//worker同Tunnel服务的心跳超时时间，通常使用默认值即可。
		HeartbeatInterval: 5* time.Second,//worker发送心跳的频率，通常使用默认值即可。
		ProcessorFactory: &tunnel.SimpleProcessFactory{
			CustomValue: "user defined interface{} value",
			ProcessFunc: exampleConsumeFunction,
			ShutdownFunc: func(ctx *tunnel.ChannelContext) {
				fmt.Println("shutdown hook",ctx.ClientId)
				daemon.Close()
			},
		},
	}

	daemon = tunnel.NewTunnelDaemon(tunnelClient,constants.DemoTableTunnelID, workConfig)
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
var counter int
func exampleConsumeFunction(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	fmt.Println("user-defined information", ctx.CustomValue)

	for _, rec := range records {
         counter++
         fmt.Println("正在消费第",counter,"行")
         fmt.Println("该行的时间戳是",rec.Timestamp)
         for _,v:=range rec.PrimaryKey.PrimaryKeys{
         	  fmt.Println(v.ColumnName,v.Value)
		 }
		for _,v:=range rec.Columns{
			fmt.Println(*v.Name,v.Value)
		}
		fmt.Println("------------------------------------------------")
	}
	fmt.Println("a round of records consumption finished") //一轮消费结束
	return nil
}
