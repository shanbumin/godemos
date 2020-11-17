package demo12

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"log"
	"os"
	"os/signal"
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/tunnel/common"
	"syscall"
)


func main() {

	//基于ots的userCheckpointer接口
	var checkpointer common.UserCheckpointer //todo implementation

	//start consume tunnel
	workConfig := &tunnel.TunnelWorkerConfig{
		ProcessorFactory: &tunnel.SimpleProcessFactory{
			CustomValue: checkpointer,
			ProcessFunc: common.ExactlyOnceIngestionFinalState,
			ShutdownFunc: func(ctx *tunnel.ChannelContext) {
				fmt.Println("shutdown hook")
			},
		},
	}

	daemon := tunnel.NewTunnelDaemon(bootstrap.TunnelClient,constants.DemoTableTunnelID, workConfig)
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


