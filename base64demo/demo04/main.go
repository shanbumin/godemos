
package main

import(
	"encoding/base64"
	"fmt"
)

func main(){
	encrypt := "Cf1WA2nBMo3H9G2UPhlLBBVBsMDl4udWr7__e6Iy93eIqLKi3EOjGhk8TkHujL1Uj6aGfZJNBzIbVE2NfNaz4pob8uiQvGaeTZdWP-8lFmAm6J1sz8N15xQkO7ADa5bNLCCqtlQbN2z7JcNenvFuID_rZGqb_1gmr-BGubGRMiMSK7RdjQYrMHaBcHLPB0UteakzcQwgKxCW7u0ECHqPJ39ne9JUG22JBWRo1ORuX5r30J_XrW3SQcdPSxfe0kvd61y12QOYh8VlOBBdBeDNnyDXefI_tDJDBFeqTXCgKu9wFkkWIZiM7WwqogaY-bvjUisbrPO4_fjJ1c0nWDOqRA"
	//为了不修改原字符串，没有直接在原字符串上追加
	encrypt = fmt.Sprint(encrypt,"==")
	_,err := base64.URLEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println(err)
	}

}
//加了两个等号，居然解码成功了，太不容易了，但是这问题也来了， 这追加的等号也不是固定的，加1个还是2个呢，如果在代码里面判断增加也很不方便啊，
//刚尝到查GO文档的好处了，带着问题找一个可以忽略最后等号的方法，看下GO文档里有没有，