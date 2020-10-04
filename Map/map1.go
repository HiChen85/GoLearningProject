package main

import (
	"fmt"
	"strconv"
)

/*
	map是一种无序的键值对集合, 可以用for循环迭代, key对应的value是事先指定好的
	在一个特定的map中,我们会指定所有key的类型和所有value的类型.
	内置的len函数可以返回map中key的数.
	因为map无序,所以每次打印map的值可能都不相同
	map是一种引用类型,所以在传参过程中是引用拷贝. 即在函数中修改map中的内容,函数外的map也会受到影响
	对于map的定义,除了使用make函数之外,还可以使用map关键字来定义
*/

func main() {
	// 定义map的第一种方式(含初始化)
	var nameIDMap = map[string]int{
		"张三": 1,
		"李四": 2,
	}
	fmt.Println(nameIDMap)
	// map定义的第二种方式
	IDNameMap := make(map[int]string, 3)
	IDNameMap[0] = "妞妞"
	IDNameMap[1] = "皮蛋"
	IDNameMap[2] = ""
	fmt.Println(IDNameMap)

	// 调用函数测试map是否会被改变
	changeMap(IDNameMap)
	fmt.Printf("%#v\n", IDNameMap)

	// map如果只声明未初始化则会定义为nil
	var tMap map[string]int
	fmt.Printf("%#v\n", tMap) // 输出结果: map[string]int(nil)

	// 使用delete函数,根据指定的key删除map中的元素,delete函数不返回任何值
	delete(IDNameMap, 2)
	fmt.Println(IDNameMap)

	// ok语法糖, 可以直接使用map_name[key] 表达式获取到key对应的value,同时,
	// 该语句会返回一个该key对应value存在与否的一个bool值
	// 如果key不存在, 则value返回该类型的0值,ok 返回false
	value, ok := IDNameMap[2]
	fmt.Printf("%#v,%#v\n", value, ok)
	fmt.Println(len(IDNameMap))

	// 单独向map中添加元素
	IDNameMap[2] = "宝宝"
	// 修改已存在的key的value
	IDNameMap[2] = "海盗船"
	fmt.Println(IDNameMap)

	// map不能用==来比较两个map是否相等,只能判断

}

// 此函数测试直接传入一个map的变量,对其修改后,原来的map是否会发生改变
// 答案是肯定的
func changeMap(m map[int]string) {
	for k, v := range m {
		// fmt.Println(k, v)
		m[k] = v + strconv.Itoa(k)
	}
}
