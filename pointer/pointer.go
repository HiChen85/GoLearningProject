package main

import "fmt"

/*
	Go语言提供了操作任何数据的指针,但是不像 C语言那样可以进行指针运算
	指针是用来存储变量地址的特殊变量, 变量地址是一个 16 进制表示的数据
*/
func main() {
	// 定义普通变量
	var a int = 10
	// 定义指针
	var aPtr *int
	// 用指针变量存储普通变量的地址
	aPtr = &a
	fmt.Println(a, aPtr)
	// 通过指针来访问对应的值, 在指针变量前再使用 * 号就可以在使用时获取到指针指向的值
	fmt.Printf("a: %#v, aPtr:%#v, *aPtr:%#v\n", a, aPtr, *aPtr)

	st := "hello"
	stPtr := &st
	fmt.Println(st, stPtr, *stPtr)
	// 修改变量的内容,看看指针的值是否被改变
	st = "世界"
	fmt.Printf("%#v, %#v\n", st, *stPtr) // 结果是指针中的值被修改了.
	// 原因在于,指针变量保存的是 st 的内存地址,如果 st 变量指向了新的值,那么对应的内存地址就会发生变化
	// 所以指针保存的 st 地址也随之发生变化. 进而用指针变量获取的内容也就是被修改后的内容
	// 通过指针修改值,是否会影响原来变量的内容?
	*stPtr = "你好"
	fmt.Printf("%#v, %#v\n", st, *stPtr) // 结果被修改了

}
