package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 3}
	fmt.Println(numIdenticalPairs(nums))

}

func game(guess, answer []int) int {
	j := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			j++
		}
	}
	return j
}

func numIdenticalPairs(nums []int) int {
	// 参考双指针法
	// 双指针法的特点在于，一次循环去完成比较，j 指针随着
	// 循环此处的增加而增加，i 指针只在 j 指针溢出数组边界
	// 时向前一位。
	var pairs int = 0
	var ptr1 = 0
	var ptr2 = ptr1 + 1
	for ptr1 < len(nums) && ptr2 < len(nums) {
		fmt.Printf("%#v, %#v\n", ptr1, ptr2)
		if nums[ptr1] == nums[ptr2] {
			pairs++
		}
		ptr2++
		if ptr2 == len(nums) {
			ptr1++
			ptr2 = ptr1 + 1
		}
	}
	return pairs
}
