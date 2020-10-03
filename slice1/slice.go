package main

import "fmt"

func main() {
	// 数组的动态初始化
	var arr1 [6]int
	fmt.Printf("%#v\n", arr1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 1
	}
	fmt.Printf("%#v\n", arr1)

	sli1 := &arr1 // 直接用指针保存数组地址，其类型是数组的指针，并不是slice类型
	fmt.Printf("%T\n", sli1)
	sli2 := arr1[3:5] // 只有进行切片操作或者直接用make去构建，才能得到切片类型的变量
	fmt.Printf("%#v\n", sli2)
	for idx, val := range sli2 {
		fmt.Println(idx, val)
	}
	fmt.Println(len(sli2), cap(sli2))
	// 对slice的修改会改变底层数组的值，因为slice类型中就包含了指向底层数组的指针
	// 因此不用再用指针去保存slice变量。
	sli2[1] = 9
	fmt.Println(arr1, sli2)
	arr1[4] = 10
	fmt.Println(arr1, sli2)

	// 向slice中追加元素使用函数append(),该函数会返回一个slice的引用，所以使用时应选择一个变量接收该引用
	arr2 := [3]int{1, 2, 3}
	// 测试向slice中增加元素，会不会改变底层数组的值
	sli3 := arr2[:]
	sli3 = append(sli3, []int{4, 5, 6, 7}...)
	fmt.Println("sli3: ", sli3, len(sli3), cap(sli3))
	fmt.Println("array2: ", arr2) // 我们发现，尽管sli3从arr2中获得，但追加新元素后，arr2的值并未改变
	sli3[0] = 99
	// 当我们追加新元素到sli3后，我们发现，再修改sli3的值，arr2的值没有发生变化，说明此时sli3底层已经指向了新的数组
	fmt.Println("sli3: ", sli3, len(sli3), cap(sli3))
	fmt.Println(arr2)
	fmt.Printf("%p, %p\n", sli3, &arr2)

	c := [5]int{1, 2, 3, 4, 5}
	sum(c[:])

	// 用make()创建一个切片
	sli4 := make([]int, 3, 6)
	fmt.Println(sli4, len(sli4))
	for i := 0; i < len(sli4); i++ {
		sli4[i] = i
	}
	sli4 = append(sli4, 4)
	fmt.Println(sli4)
	fmt.Printf("%p\n", sli4)
	sli4 = append(sli4, []int{5, 6, 7}...)
	fmt.Println(sli4)
	fmt.Printf("%p\n", sli4)

	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s1, s2)

	var arr3 = [...]int{10, 20, 30, 40}
	for idx := range arr3 {
		arr3[idx] *= 2
	}
	fmt.Println(arr3)

	// 测试copy函数, copy函数能将一个silce拷贝到一个新的slice中. 并返回拷贝元素的个数
	newSli := make([]int, 5, 10)
	oldSli := []int{1, 3, 5, 7, 9}
	num := copy(newSli, oldSli)
	fmt.Println(num)
	fmt.Printf("%#v\n", newSli)

}

// 向函数传入切片
// 当我们拥有一个函数需要传入一个数组的时候，我们可能总是需要把参数声明为切片,然后将数组进行切片操作
// 这一步的操作等于是在为数组创建一个引用，可以更加方便的操作数组，直到该切片扩容，自动指向新数组为止
func sum(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	fmt.Println(sum)
	return
}
