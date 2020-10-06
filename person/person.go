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
