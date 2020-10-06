package main

import "fmt"

func main() {
	reversePrint(5)
}

// 递归倒序打印从n开始的数字
func reversePrint(n int) int {
	if n == 1 {
		fmt.Println(1)
		return 1
	}
	fmt.Println(n)
	return reversePrint(n - 1)
}
