package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
	高阶函数于函数式编程:
	1. 函数名()表示函数调用.
	2. 函数名本身可以理解为指向函数体的内存地址
	3. 函数可以作为另一个函数的参数, 也可以作为某一个函数的返回值.这样的函数称之为高阶函数
*/

// 函数作为参数的例子
func calcu(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

//函数作为返回值的例子
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		return nil, errors.New("无法识别操作符")
	}

}

func main() {
	fmt.Println(calcu(1, 2, add))

	// 我们可以利用函数作为返回值的特点,定义一些操作.然后用变量保存该操作后便于计算
	ad, err := do("+")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ad(1, 9))

	fibo := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println("fibo:", fibo())
	}

	suf := makeSuffix(".txt")
	fmt.Println(suf("123"))
	fmt.Println(suf("123.jpg"))
}

// 匿名函数: 不指定函数名的都称为匿名函数
// 在go语言的函数体内部定义函数时,只能定义匿名函数
// 匿名函数需要被某个变量接收或者作为立即执行的函数
// 匿名函数多用于实现匿名函数和闭包

// 闭包
// 闭包 = 函数 + 引用环境
// 闭包: 一个外部函数的局部变量可以进入到函数内部的匿名函数中进行计算.
// 内部函数在计算完成后返回外部函数的局部变量,而外部函数则返回整个匿名函数
// 返回的函数需要由某个变量接收,以便在其他地方调用. 返回的这个匿名函数和匿名函数所
// 返回的局部变量一起构成一个闭包.用来保存这个局部变量的值.或者是某种现场.
// 但是通常,把返回的匿名函数称为闭包
// 闭包的外部引用环境(外部函数的局部变量)可根据需要在内层函数(闭包函数)返回.
// 即不一定每次都要返回
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包示例1
func fibonacci() func() int {
	x, y := 1, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

//闭包示例2
// 此示例使用了外部函数的局部变量,但并未修改该变量.
// 仅仅是使用了外部变量的值,但是这个值在被传入后就
// 一直被保存,知道闭包结束调用
func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		switch {
		case !strings.Contains(fileName, ".") && !strings.HasSuffix(fileName, suffix):
			fileName += suffix
			return fileName
		case strings.Contains(fileName, ".") && !strings.HasSuffix(fileName, suffix):
			fmt.Println("文件格式已存在", strings.Split(fileName, ".")[1])
			return fileName
		default:
			return ""
		}
	}
}
