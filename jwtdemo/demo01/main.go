package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//构建和签名令牌的简单示例
func main() {

	//1.自定义一个secret
	hmacSampleSecret:="qTw32gOomGLMAatO2HcmGp2hsq8P3cb7"

	//2.声明想传递的一些声明
	user:= map[string]interface{}{}
	user["name"] = "sam"
	user["age"] = 18
	user["sex"] ="male"
	dataByte,_:= json.Marshal(user)
	var dataStr = string(dataByte)

	//3.使用Claim保存json 这里是个例子，并包含了一个故意签发一个已过期的token
	//data := jwt.StandardClaims{
	//	 Subject:dataStr,
	//	 ExpiresAt:time.Now().Unix()-1000}
	//todo 您也可以使用非标准Claims，自己直接使用MapClaims也可以的额

	data:= jwt.MapClaims{
		"sub":dataStr,
		"exp":time.Now().Unix() + 600,
		"name":"sam",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,data)

	//4.生成token串
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	fmt.Println(tokenString, err)




}
