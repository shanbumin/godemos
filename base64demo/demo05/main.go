package main

import(
	"encoding/base64"
	"fmt"
)

func main(){
	encrypt := "Cf1WA2nBMo3H9G2UPhlLBBVBsMDl4udWr7__e6Iy93eIqLKi3EOjGhk8TkHujL1Uj6aGfZJNBzIbVE2NfNaz4pob8uiQvGaeTZdWP-8lFmAm6J1sz8N15xQkO7ADa5bNLCCqtlQbN2z7JcNenvFuID_rZGqb_1gmr-BGubGRMiMSK7RdjQYrMHaBcHLPB0UteakzcQwgKxCW7u0ECHqPJ39ne9JUG22JBWRo1ORuX5r30J_XrW3SQcdPSxfe0kvd61y12QOYh8VlOBBdBeDNnyDXefI_tDJDBFeqTXCgKu9wFkkWIZiM7WwqogaY-bvjUisbrPO4_fjJ1c0nWDOqRA"
	_,err := base64.RawURLEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println(err)
	}
}

//总结
//base64编码过程有两部特殊操作
//url safe 将+/字符串转化成_-
//no padding is add  末尾不增加=号