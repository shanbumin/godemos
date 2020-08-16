package main

import (
	"fmt"
)


func double(x *int) {
	//todo 这一句把 x 指向的值（也就是 &a 指向的值，即变量 a）变为原来的 2 倍。但是对 x 本身（一个指针）的操作却不会影响外层的 a，所以下面 x=nil掀不起任何大风大浪。
	*x += *x
	//todo 这得稍微思考一下，才能得出这一行代码根本不影响的结论。因为是值传递，所以 x 也只是对 &a 的一个拷贝。所以这里仅仅是将一个指针拷贝置空而已
	x = nil
}


func main() {
	var a = 3
	double(&a)
	fmt.Println(a) // 6

	fmt.Println("---------------------")

	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false

}