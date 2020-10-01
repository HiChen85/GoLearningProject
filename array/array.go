package main

import "fmt"

// 数组在 go 中是值类型.传参时,如果不传入数组引用,则会发生数组的拷贝.
// 数组的长度也是该数据类型的一部分,所以长度不同,元素类型相同的数组是两个变量.彼此不相同

func main() {
	// 完整数组的声明
	var array [5]int
	// 带有初始化的声明
	// 初始化过程中,如果数组元素没达到数组长度,其余的值被置为 0
	var array1 [5]int = [5]int{1, 2, 3, 4} // 最后一个元素被置为 0
	fmt.Println(array, array1)
	// 由于 go 可以推断数据类型,因此可以按照如下方式来定义数组
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// 在函数体内可以使用更加简单的数组定义方式
	array3 := [5]int{1, 3, 4, 5, 7}
	fmt.Println(array3)

	// 数组是可变的: 即 可以使用数组名[i]的方式对数组中的元素进行更改
	array3[1] = 9
	fmt.Println(array3)

	// 数组的普通值传递会发生值拷贝，内部操作不会修改原始数组的值
	arr := [...]int{1, 2, 3}
	valueCopy(arr)
	fmt.Println(arr)
	// 传递引用后数组值发生变化
	refCopy(&arr)
	fmt.Println(arr)
}

func valueCopy(arr [3]int) {
	for idx, val := range arr {
		arr[idx] = val + 1
	}
	fmt.Println(arr)
}

func refCopy(arr *[3]int) {
	for idx := range arr {
		arr[idx]++
	}
	fmt.Println(arr)
}
