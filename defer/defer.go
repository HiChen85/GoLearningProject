package main

import "fmt"

// defer是go语言的一种延迟处理机制
// 该机制遵循后进先出,也就是后注册的延迟函数先执行,先注册的后执行规则
// 由于defer的延迟处理机制,一般讲defer用于资源的关闭,资源清理等场景

// go语言的return操作分为两部分,给返回值赋值和RET操作,defer语句执行的操作
// 就在给返回值赋值之后

// defer注册延迟函数时,被注册的函数内的所有参数都要是确定值,也就是说如果传参是一个函数调用,则该函数会先于defer执行
// 如果是一个变量,则其之前定义的值就被传入,即便之后被修改,仍然是被修改前的变量值

func main() {
	// deferStart()
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

// A, 1, 2, 3
// B, 10, 2, 12
// BB, 10 ,12, 22
// AA, 1, 3, 13

/*
执行结果:
	func start...
	func end...
	3
	2
	1
*/
func deferStart() {
	fmt.Println("func start...")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("func end...")
}

func f1() int {
	x := 1
	defer func() {
		x++
	}()
	return x // return语句分为两步,给返回值赋值和RET操作

	// 这里虽然return x, 但其实x的值在return操作中已经赋值了,所以赋值之后的defer操作对x的修改只是修改了函数内的局部变量
	// 对返回值并无影响
}

// 此函数不同于上一个函数,我们在函数定义时指定了函数的返回值名称,因此对此x的修改会作用到返回值中
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5

	//这里, return操作时,先将5赋值给x,然后发现注册了defer,所以执行defer,函数操作了定义好的返回值x
	// 所以return语句中的返回值发生了变化.
}

func f3() (y int) {
	x := 1
	defer func() {
		x++
	}()
	return x

	// 此函数在return操作中,将x的值赋值给定义好的返回值y,然后执行defer操作修改x的值
	// 但x并没有被最终返回,所以最终结果并没有被改变
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5

	//在这里函数中注册defer函数,内部形参只是名字和外部定义的返回值名字相同,但根据函数作用域的不同以及函数值传递的
	// 特点,最终返回结果仍然不会发生变化
}

// defer面试题
func calc(index string, a, b int) int {
	rst := a + b
	fmt.Println(index, a, b, rst)
	return rst
}
