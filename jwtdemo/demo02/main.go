package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//解析和验证令牌的简单示例
func main() {


	//token字符串
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYW1lIjoic2FtIiwibmJmIjoxNDQ0NDc4NDAwfQ.r9akkIdQg6eRdqZ-alF1piHn8yPfBbNGrHSNwJaDIjw"
	hmacSampleSecret:=[]byte("qTw32gOomGLMAatO2HcmGp2hsq8P3cb7")

    //解析
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["name"])
	} else {
		fmt.Println(err)
	}




	
}
