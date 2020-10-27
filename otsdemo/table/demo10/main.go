package main

import (
	"otsdemo/sample"
	"otsdemo/start"
)

//原子计数器
//将列当成一个原子计数器使用，对该列进行原子计数操作，可用于为某些在线应用提供实时统计功能，例如统计帖子的PV（实时浏览量）等。

func main() {

	   sample.UpdateRowWithIncrementColumn(start.Client,sample.TableConditionName)
}
