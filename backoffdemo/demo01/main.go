package main

import (
	"errors"
	"github.com/cenkalti/backoff"
	"log"
	"time"
)


func test()error{
	 return  errors.New("again")
}
//使用重试功能重试可能失败的操作
func main() {

	jitter:=0.33
	multiplier:=2.0
	interval:= 50 * time.Millisecond //每次等待的最小值
	maxInterval:= 3* time.Second  //每次等待的最大值 todo 并不是严格的，只是在3附近徘徊
	maxElapsed:=  75 * time.Second
	b := backoff.NewExponentialBackOff()
	b.RandomizationFactor = jitter //随机因素 0.33
	b.Multiplier = multiplier //乘数 2.0
	b.InitialInterval = interval //初始化间隔 0.016s
	b.MaxInterval = maxInterval
	b.MaxElapsedTime = maxElapsed
	b.Reset()

	for {
		err:=test()
		if err == nil {
			break
		} else {
			nextBkoff := b.NextBackOff()
			log.Println("nextBkoff=",nextBkoff)
			if nextBkoff == backoff.Stop {
				log.Println("结束了")
				return
			}
			time.Sleep(nextBkoff)
		}
	}






}
//todo 重试区间是75s
//2020/11/21 19:38:42 nextBkoff= 17.105213ms
//2020/11/21 19:38:42 nextBkoff= 41.303552ms
//2020/11/21 19:38:42 nextBkoff= 70.951017ms
//2020/11/21 19:38:42 nextBkoff= 122.738094ms
//2020/11/21 19:38:42 nextBkoff= 243.266751ms
//2020/11/21 19:38:43 nextBkoff= 575.131253ms
//2020/11/21 19:38:43 nextBkoff= 713.320432ms
//....
//2020/11/21 19:39:56 nextBkoff= 1.034164932s
//2020/11/21 19:39:57 nextBkoff= 1.168843515s
//2020/11/21 19:39:58 nextBkoff= -1ns
//2020/11/21 19:39:58 结束了