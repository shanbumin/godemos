package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//构建和签名令牌的简单示例
func main() {

	hmacSampleSecret:="qTw32gOomGLMAatO2HcmGp2hsq8P3cb7"

	//创建一个新的令牌对象，指定签名方法和您希望包含的声明。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"name":"sam",
	})

	//生成token串
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	fmt.Println(tokenString, err)




}
