package main

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tunnel"
	"net/http"
	"otsdemo/constants"
	"time"
)

//todo tunnel client
//初始化tunnel client时可以通过NewTunnelClientWithConfig接口自定义客户端配置，使用不指定config初始化接口或者config为nil时会使用DefaultTunnelConfig。
func main() {


    conf:=&tunnel.TunnelConfig{
		MaxRetryElapsedTime: 45 * time.Second,//最大指数退避重试时间。
		RequestTimeout:      30 * time.Second, //HTTP请求超时时间。
		Transport:           http.DefaultTransport, //http.DefaultTransport。
	}
	tunnel.NewTunnelClientWithConfig(constants.EndPoint,constants.InstanceName,constants.AccessKeyId,constants.AccessKeySecret,conf)
	
}
