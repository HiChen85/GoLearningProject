package main

/*
	值类型: 基本数据类型都属于值类型. 即 任何的传参过程都会拷贝一份改变量的副本到函数内,函数内部的任何操作不会修改原本变量的内容
	引用类型: 使用引用类型传参,函数内的操作会影响该变量的值, 所有的引用类型统一使用make函数进行内存分配
	golang中的值类型: 基本数据类型, array, 结构体变量. 因此,在程序中如果想通过某函数修改这些变量的内部字段或者值,则需要用对应的指针类型
	保存这些变量的地址然后进行操作. 使用等号的过程就是所谓值拷贝的过程
	golang中的引用类型: slice, map, interface, channel都是引用类型, 除了赋值过程,任何的操作都会修改这些变量的内容.我们在使用时无需再进行
	额外的指针赋值操作,可以直接使用这些变量的变量名称
*/

import (
	"fmt"
)

var X = "G"

func main() {
	/*
		1. 变量声明方式: var identifier [type] = expression(or value)
	*/

	// basic data type
	var a bool = true           // default value of bool is false
	var b int = 15              // default value of int is 0
	var c float64 = 3.141592653 // default value of float64 is 0
	var d string                // // default value of string is ""

	fmt.Printf("%#v, %#v, %#v, %#v\n", a, b, c, d)
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	/*
		2. 简洁变量命名法: 只适用于函数内部: x := expression (or value)
	*/
	f := "hello world"
	fmt.Println(f)

	n()
	m()
	n()

	/*
		在go语言中,如果函数没有接收参数,但是内部用到了某个变量,函数会先寻找其局部变量,如果没有局部变量,则直接向全局变量寻找.
	*/
}

func n() {
	fmt.Println(X)
}
func m() {
	X = "O"
	fmt.Println(X)
}
