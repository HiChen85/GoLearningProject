package main

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05 Monday"

//在包前用简易名称代替包名

/*
	time 包为我们提供了一个类型. time.Time 类型作为值使用.包含测量时间日期以及显示日期时间的函数

*/

func main() {
	// 获取当前时间
	t := time.Now()
	fmt.Println(t)
	// 根据获得的 time 对象可以获取当前的部分日期或者时间信息
	fmt.Println(t.Day())
	fmt.Println(t.Minute())
	// 日期格式化: Format()函数会根据一个指定的格式化日期格式来输出日期和时间
	st := t.Format(timeFormat)
	fmt.Println(st)
}
