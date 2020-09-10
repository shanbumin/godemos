package main

import "fmt"



const (
	DefaultJobGroup = "default"
	KindCommon   = iota //执行间隔到了，允许执行的每一台节点上都会运行
	KindAlone    // 执行间隔到了，只允许在某一台节点上运行。(在执行没有结束前，哪怕下一个执行间隔到了也将忽略)
	KindInterval
)



func main() {
	fmt.Println(KindCommon)
	fmt.Println(KindAlone)
	fmt.Println(KindInterval)
}
