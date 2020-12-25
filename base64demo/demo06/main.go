package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str:="sam is a good man."
    src:=[]byte(str)
	// 计算加密后数据的长度
	n := base64.StdEncoding.EncodedLen(len(src))
	dst := make([]byte,n) // 创建容器



	base64.StdEncoding.Encode(dst,src) // 加密数据
	fmt.Printf("%s\n",dst) // 转换为字符串： YWJjZGVmZw==

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
}
