package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//解析和验证令牌的简单示例
//@link https://www.jsonwebtoken.io/
func main() {


	//token字符串
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTc3MzA5MjksIm5hbWUiOiJzYW0iLCJzdWIiOiJ7XCJhZ2VcIjoxOCxcIm5hbWVcIjpcInNhbVwiLFwic2V4XCI6XCJtYWxlXCJ9In0.iPZ2KJqyPPi2vObpVSxbq2CiUacY_4EFzxmmmEHIGqc"
	hmacSampleSecret:=[]byte("qTw32gOomGLMAatO2HcmGp2hsq8P3cb7")

	///将token字符串转换为token结构体
	//todo 注意如果token的生成设置过期时间之后，则这里会主动检查token的过期情况的额
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//}
		return hmacSampleSecret, nil
	})
	fmt.Println(err) //如果过期的话，这里已经打出来了

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//todo 这里没有必要去主动验证是否过期了
		////校验下token是否过期
		//succ := claims.VerifyExpiresAt(time.Now().Unix(),true)
		//fmt.Println("succ",succ)
		//获取token中保存的用户信息
        //fmt.Println(claims["sub"])
		fmt.Printf("%#v\r\n",claims)
	} else {
		fmt.Println(err)
	}



}
