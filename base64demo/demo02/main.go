package main

import(
	"encoding/base64"
	"fmt"
)

func main(){
	encrypt := "Cf1WA2nBMo3H9G2UPhlLBBVBsMDl4udWr7__e6Iy93eIqLKi3EOjGhk8TkHujL1Uj6aGfZJNBzIbVE2NfNaz4pob8uiQvGaeTZdWP-8lFmAm6J1sz8N15xQkO7ADa5bNLCCqtlQbN2z7JcNenvFuID_rZGqb_1gmr-BGubGRMiMSK7RdjQYrMHaBcHLPB0UteakzcQwgKxCW7u0ECHqPJ39ne9JUG22JBWRo1ORuX5r30J_XrW3SQcdPSxfe0kvd61y12QOYh8VlOBBdBeDNnyDXefI_tDJDBFeqTXCgKu9wFkkWIZiM7WwqogaY-bvjUisbrPO4_fjJ1c0nWDOqRA"
	_,err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println(err)
	}
}
//执行结果
//illegal base64 data at input byte 34
//解码报错了，根据提示意思大概能猜到是第34位的字符‘_’ base64不认识，去查了下base64的索引表


//发现base64的字符集内没有我们第34位对应的字符‘_’,当时想是不是做了类似urlencode的编码，防止http传输过程中部分字符转义，继续查base64文档发现了 如下一段内容

//可以确定是将+和/分别改成了-和_, 那我们应该做一下字符串替换把-和_改回来就行，本来想直接在代码里面写字符串替换，后来想GO标准库还没细看，里面是不是还有现成的方法之前没发现，查了下标准库发现
//todo URLEncoding和我们刚用的StdEncoding有一些区别，是用于URL和文件名，刚我们碰到的问题也是URL问题