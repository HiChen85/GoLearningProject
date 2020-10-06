package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", alterableParemeterFunc(1, 2, 3, 4, 5))
	fmt.Println(recursion(4))
	fmt.Println(fibonacci(4))

	var cal calculation = add
	fmt.Println((cal(1, 2)))
}

// 接收可变参数的函数定义法
// 可变参数的意思就是接收的参数不止一个,但不知具体是几个
// 所有的可变参数都由一个变量管控,在函数内部,所有的可变参数都按照可变参数名作slice处理
// 需要注意的是,如果定义可变参数,则可变参数的类型必须统一
func alterableParemeterFunc(alter ...int) []int {
	for i := 0; i < len(alter); i++ {
		alter[i] = i + 1
	}
	return alter
}

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	return recursion(n-1) * n
}

// n >= 1
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 定义函数类型
type calculation func(int, int) int

// 定义两个符合上述类型的函数
// 当我们定义了函数类型后,就可以创建该类型的变量,然后用该变量去保存函数.
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
