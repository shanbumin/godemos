package main

import (
	"encoding/base64"
	"fmt"
	"log"
)


//TODO 如果str中没有特殊的+/字符，且末尾无补位=现象，三者的加解密结果都是一样，彼此之间可以相互解密对方的加密手法
//todo 我们去下面的案例中体会三者的差异，保险起见使用第三种是最包罗万象的额
func main() {

   //法一
	var str = "sam is a good man."
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("base64标准加密的结果:",encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("base64标准解密的结果:", string(decoded))
	fmt.Println("-------------------------------------------------------------")
	//法二
	encoded2 := base64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println("base64URL加密的结果:",encoded2)
	decoded2, err := base64.URLEncoding.DecodeString(encoded2)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("base64UR解密解密的结果:", string(decoded2))
	fmt.Println("-------------------------------------------------------------")
	//法三
	encoded3 :=base64.RawURLEncoding.EncodeToString([]byte(str))
	fmt.Println("base64RawURL加密的结果:",encoded3)
	decoded3,err:=base64.RawURLEncoding.DecodeString(encoded3)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("base64RawUR解密的结果:", string(decoded3))



}
