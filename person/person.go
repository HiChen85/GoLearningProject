package person

import "fmt"

// Person 人的定义
type Person struct {
	name string
	age  int
}

// SayHello Person's action
// go语言方法的定义,将所谓的this指针放在了函数名之前,将其称之为接受者
// 效果同this指针一样,都是指一个特定的自定义类型所具备的方法
func (p *Person) SayHello() {
	fmt.Println("Hello, my name is", p.name)
	fmt.Println("I'm", p.age, "year-old")
}

// SetName 设置对象的名称
func (p *Person) SetName(name string) {
	p.name = name
}

// SetAge 设置年龄
func (p *Person) SetAge(age int) {
	p.age = age
}

// 使用工厂方法 + 类型的访问权限来创建类型的实例
// 工厂方法一般使用New开头, 返回一个指定类型的一个指针
// 工厂方法类似于java中的构造器

// NewPerson 工厂方法创建Person实例
func NewPerson(name string, age int) *Person {
	p := new(Person)
	p.name = name
	p.age = age
	return p
}
