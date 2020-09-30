package main

import (
	"fmt"
)

/*
	GO对值之间的比较有着严格的限制,只有相同值之间才允许进行比较.
*/

func boolComparison() {
	a := 10
	if a == 5 {
		fmt.Println(a)
	} else {
		fmt.Println(-a)
	}
	b := true
	if !b {
		fmt.Println("b不是真的")
	} else {
		fmt.Println("b是真的")
	}
}

/*
	Go有基于架构的基本数据类型: int, uint, uintptr. 这些类型的长度是根据程序所运行的操作系统决定的.
	int和uint在32为操作系统上占32位(8字节), 在64为系统上占64位(8字节), uintptr被设定为能够存放一个指针即可.
	注意: int是计算最快的一种类型
	与操作系统无关的数据类型从类型名称即可看出来: int8, int16, int32, int64.
	float32(精确到小数点后7位), float64(精确到小数点后15位).
	在go中应尽可能的使用float64,因为math包中所有的数据都是基于float64进行设计的
*/

/*
	在数字前加0,例如077来表示8进制数, 用0x来表示16进制数.用e表示10的n次幂.
	在go中不同类型之间的赋值,需要强制转换
*/

func intOperation() {
	var a int = 100
	var b int64 = 10
	var c int64
	c = int64(a) + b // 不强制转换则出错
	fmt.Println(c)
}

//digital operation
/*
	位运算只适用于整数,且需要他们具备等长位模式.
	二元位运算:
	& 按位与: 同位置上 1&1=1, 1&0=0, 0&0=0. 类似于逻辑运算&&.
	| 按位或: 同位置上 1|1=1, 1|0=1, 0|0=0.
	^ 按位异或: 同位置上 1^1=0, 0^0=0, 1^0=1.
	一元位运算:
	^ 作为一元运算符,表示按位取反,对有符号整数,最高位的数字表示符号位,也要按位取反.
	<< 左移运算符: 对数字的二进制形式进行向左移动,右边空白补0, 移动的位数表示前面的数字乘2的n次方
	>> 右移运算符: 对数字的二进制形式进行向右移动,左边空白补0, 移动的位数表示前面数字除以2的n次方

	运算符优先级:
	第一优先级: *, /, %, <<, >>, &, &^
	第二优先级: +, -, |, ^
*/

func sumOfTwo(a, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return sumOfTwo(sum, carry)
}

func main() {
	// boolComparison()
	// fmt.Printf("%#v\n", ^2)
	// fmt.Printf("%#v\n", 1<<2)
	// fmt.Printf("%#v\n", 8>>2)
}
