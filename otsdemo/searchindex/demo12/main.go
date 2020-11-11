package main

import (
	"otsdemo/bootstrap"
	"otsdemo/constants"
	"otsdemo/searchindex/servers"
)


//通配符查询中，要匹配的值可以是一个带有通配符的字符串，目前支持星号（*）和问号（?）两种通配符。
//要匹配的值中可以用星号（*）代表任意字符序列，或者用问号（?）代表任意单个字符，且支持以星号（*）或问号（?）开头。例如查询“table*e”，可以匹配到“tablestore”。
func main() {
	servers.WildcardQuery(bootstrap.Client,constants.DemoTable,constants.DemoTableIndex)
}
