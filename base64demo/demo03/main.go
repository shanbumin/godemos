package main

import(
	"encoding/base64"
	"fmt"
)

func main(){
	encrypt := "Cf1WA2nBMo3H9G2UPhlLBBVBsMDl4udWr7__e6Iy93eIqLKi3EOjGhk8TkHujL1Uj6aGfZJNBzIbVE2NfNaz4pob8uiQvGaeTZdWP-8lFmAm6J1sz8N15xQkO7ADa5bNLCCqtlQbN2z7JcNenvFuID_rZGqb_1gmr-BGubGRMiMSK7RdjQYrMHaBcHLPB0UteakzcQwgKxCW7u0ECHqPJ39ne9JUG22JBWRo1ORuX5r30J_XrW3SQcdPSxfe0kvd61y12QOYh8VlOBBdBeDNnyDXefI_tDJDBFeqTXCgKu9wFkkWIZiM7WwqogaY-bvjUisbrPO4_fjJ1c0nWDOqRA"
	_,err := base64.URLEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println(err)
	}
}


//执行一下 又报错了，心累，看了下错误信息illegal base64 data at input byte 340，跟第一步报错的大致一样只是最后的位置变到了340，
//说明我们前面转义的问题用这个方法还是解决了，但是执行到最后的时候又有不标准字符了，继续查base64的文档，又有新发现，

//大致意思是，如果编码的时候字节不足会在最后加一到两个=号，但看我们的字符串最后没有=，解码的时候解到最后又报错了，赶紧给字符串手动加个等号试试