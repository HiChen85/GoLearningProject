package main

import "fmt"

func main() {
	// 切片的扩容策略: 按照原有切片的容量进行2倍扩容
	sli1 := make([]int, 5) // 默认不指定cap时, cap = len
	for i := 0; i < len(sli1); i++ {
		sli1[i] = i + 1
	}
	fmt.Printf("%p, %#v\n", sli1, sli1)
	fmt.Println(len(sli1), cap(sli1))
	for i := 0; i < 5; i++ {
		sli1 = append(sli1, i+6)
	}
	fmt.Printf("%p, %#v\n", sli1, sli1)
	fmt.Println(len(sli1), cap(sli1))
	for i := 0; i < 5; i++ {
		sli1 = append(sli1, i)
	}
	fmt.Printf("%p, %#v\n", sli1, sli1)
	fmt.Println(len(sli1), cap(sli1))
	// 从这里可以看出, 当扩容后,原有的slice指向新的底层数组
	// 对slice新的修改不会再影响原有的slice
}
