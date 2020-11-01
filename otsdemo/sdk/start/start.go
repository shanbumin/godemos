package start

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/v5/tablestore"
	"otsdemo/sample"
)

var Client *tablestore.TableStoreClient




func init()  {
	//初始化``TableStoreClient``实例。
	//endPoint是表格存储服务的地址（例如'https://instance.cn-hangzhou.ots.aliyun.com:80'），必须以'https://'开头。
	//accessKeyId是访问表格存储服务的AccessKeyID，通过官方网站申请或通过管理员获取。
	//accessKeySecret是访问表格存储服务的AccessKeySecret，通过官方网站申请或通过管理员获取。
	//instanceName是要访问的实例名，通过官方网站控制台创建或通过管理员获取。
	Client=tablestore.NewClient(sample.EndPoint, sample.InstanceName,sample.AccessKeyId, sample.AccessKeySecret)
}
