package main

import (
	"fmt"
	"unsafe"
)

//Len，cap 的转换流程如下：
//Len: &s => pointer => uintptr => pointer => *int => int
//Cap: &s => pointer => uintptr => pointer => *int => int


//todo  此处的uintptr(8)  uintptr(16)是个什么鬼东西呢？
func main() {



	s := make([]int, 9, 20)
    var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
    fmt.Println(Len, len(s)) // 9 9



    var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
    fmt.Println(Cap, cap(s)) // 20 20
}


