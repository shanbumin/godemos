package main
import "fmt"


func double(x int) {
	x += x
}

//为什么需要指针类型呢？
func main() {


	var a = 3
	double(a)
	fmt.Println(a) // 3
	//非常简单，我想在 double 函数里将 a 翻倍，但是例子中的函数却做不到。为什么？因为 Go 语言的函数传参都是 值传递。double 函数里的 x 只是实参 a 的一个拷贝，在函数内部对 x 的操作不能反馈到实参 a。
	//如果这时，有一个指针就可以解决问题了！这也是我们常用的“伎俩”。


}