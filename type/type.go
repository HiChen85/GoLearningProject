package main

import "fmt"

// 类型声明的语句: type 类型名称 底层类型.

// Celsius 声明一个摄氏度类型,它和float64具有相同的底层类型, 即float64类型
type Celsius float64 // 这种结构都叫做自定义类型

// Fehrenheit 华氏温度 和float64具有相同的底层结构
type Fehrenheit float64

const (
	// AbsoluteZeroC 绝对0度
	AbsoluteZeroC Celsius = -273.15
	// FreezingC 结冰温度
	FreezingC Celsius = 0
	// BoilingC 开水温度
	BoilingC Celsius = 100
)

// CToF 摄氏温度转华氏温度
func CToF(c Celsius) Fehrenheit {
	return Fehrenheit(c*9/5 + 32)
}

// FToC 华氏度转摄氏度
func FToC(f Fehrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

/*
	此处我们声明了两种不同的数据类型,尽管他们拥有相同的底层结构. 他们不可以被相互比较, 因此如需比较操作,则需对其进行显示的类型转换
	才能进行值的比较操作. 对于每一个类型来说,都有一个T()操作,来对和T类型具有同样底层结构的其他类型进行显示的类型转换,但是拥有相同
	底层结构的类型之间可能存在语义的差别,所以尽管可以进行显示类型转换,但可能会存在语义上的错误.
*/

func main() {
	var tempreature Celsius = 32
	fmt.Println(CToF(tempreature))
}
