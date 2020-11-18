package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"os"
	"os/signal"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/tunnel/show/conf"
	"strconv"
	"syscall"
	"time"
)






func main() {

	log.Println("Waiting...")

	//创建客户端
	tunnelClient:=bootstrap.DefTunnelClient
	var daemon *tunnel.TunnelWorkerDaemon
	//start consume tunnel
	workConfig := &tunnel.TunnelWorkerConfig{
		HeartbeatTimeout: 10 * time.Second,//worker同Tunnel服务的心跳超时时间，通常使用默认值即可。
		HeartbeatInterval: 5* time.Second,//worker发送心跳的频率，通常使用默认值即可。
		ProcessorFactory: &tunnel.SimpleProcessFactory{
			//CustomValue: "user defined interface{} value",
			ProcessFunc: exampleConsumeFunction,
			ShutdownFunc: func(ctx *tunnel.ChannelContext) {
				fmt.Println("shutdown hook",ctx.ClientId)
				daemon.Close()
			},
		},
		LogConfig:&conf.DefaultLogConfig,
		LogWriteSyncer:conf.DefaultSyncer,
		BackoffConfig:&conf.DefaultBackoffConfig,
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

var count int64

func exampleConsumeFunction(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	fmt.Println("user-defined information", ctx.CustomValue)

	for _, rec := range records {
		//开场白
		count++
		welcome(rec.Timestamp,count)

         for _,v:=range rec.PrimaryKey.PrimaryKeys{
         	  fmt.Println(v.ColumnName,v.Value)
		 }
		for _,v:=range rec.Columns{
			fmt.Println(*v.Name,v.Value)
		}
	}
	fmt.Println("a round of records consumption finished") //一轮消费结束
	return nil
}

func welcome(timestamp int64,count int64){
	var tell string
	if timestamp <=0{
		tell="[ BaseData without timestamp ]"
	}else{
		sec, _ := strconv.ParseInt(Substr(strconv.Itoa(int(timestamp)), 0, 10), 10, 64)
		tell=time.Unix(sec, 0).Format("2006-01-02 15:04:05")
	}
	log.Println(">>>>>  ",tell,"Record:",count)
}


//从string中index位置开始，截取n个字节长度
func Substr(s string, index int, n int) string {
	L := len(s)
	if index < 0 || index >= L || s == "" {
		return ""
	}
	end := index + n
	if end >= L {
		end = L
	}
	if end <= index {
		return ""
	}
	return s[index:end]
}


