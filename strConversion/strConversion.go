package main

import (
	"fmt"
	"strconv"
)

/*
	strconv 包提供字符串和其他数据类型之间转换操作的函数
	该包包含了一些用于获取程序运行的操作系统环境下 int 所占的位数
	例如: strconv.IntSize 变量
	任意类型 T 转换为字符串总是可以进行的
*/

func main() {
	fmt.Println(strconv.IntSize) //64为操作系统结果为 64
	//数字转换为对应字符串的操作
	// strconv.Itoa(i int) string 函数,接收一个整数,返回其字面表示的数字的字符串形式
	fmt.Printf("%#v\n", strconv.Itoa(100))

	// 字符串转为int
	val1, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换出错")
	}
	fmt.Printf("%#v\n", val1)

	//字符串转 int64
	numInt64, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println("转换出错")
	} else {
		fmt.Printf("%#v\n", numInt64)
	}

	// 字符串转 float32
	numFloat32, err := strconv.ParseFloat("3.14159", 32)
	if err != nil {
		fmt.Println("转换出错")
	} else {
		fmt.Printf("%#v\n", numFloat32)
	}
	//  字符串转 float64
	numFloat64, err := strconv.ParseFloat("2.71234", 64)
	if err != nil {
		fmt.Println("转换出错")
	} else {
		fmt.Printf("%#v\n", numFloat64)
	}
}
