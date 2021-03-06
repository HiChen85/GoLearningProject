package main

import "GoLearningProject/person"

// 在type目录下,有关于自定义类型的定义
// 即 创建一些基本数据类型作为底层类型的新类型,它们除了有基本数据类型的特点之外
// 属于一个全新类型.可以根据底层类型的特点进行相应的操作,同样可以为这个类型添加
// 新的操作.

// 类型别名格式: type TypeAlias = 基本类型
// 这种结构只是对原有的类型起了另一个名字,它们同属一个类型.

// 自定义类型和类型别名的却别:
// 当我们用自定义类型定义一个基本数据类型,打印类型信息后会显示新的类型名
// 当用类型别名指定数据类型,打印类型时仍然是基本数据类型

// structure 结构体,由一些字段名和对应的类型构成的集合

func main() {
	p := new(person.Person) // 获取类型实例的指针, 但是在go中有个语法糖,就是结构体的实例和实例指针,都可以直接用 变量名.字段名访问
	p.SetName("Chen")
	p.SetAge(18)
	p.SayHello()
}
