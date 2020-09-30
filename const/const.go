package main

import "fmt"

/*
	go语言中的常量是用const关键字修饰, 一般定义在函数体外部,便于查看
	go语言中有一个特殊常量iota, 每当遇到一次const关键字,iota就会重置为0
	在一个const修饰的常量内部,每新增一个常量,iota会+1

*/

// 圆括号内的多行赋值是不需要逗号的
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
