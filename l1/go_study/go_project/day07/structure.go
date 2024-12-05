package day07

import "fmt"

// 结构体是一个复合类型，用于表示一组数据
// 结构体由一系列属性组成，每个属性都有自己的类型和值
// 定义结构体
type Person struct {
	name  string
	age   int
	hobby []string
}

func Structure() {
	//先后顺序
	var p1 = Person{"luomanfei", 19, []string{"篮球", "足球"}}
	fmt.Println(p1.name, p1.age, p1.hobby)
	//关键字
	var p2 = Person{name: "luomanfei", age: 19, hobby: []string{"篮球", "足球"}}
	fmt.Println(p2.name, p2.age, p2.hobby)
	//先声明再赋值
	var p3 Person
	p3.name = "luomanfei"
	p3.age = 19
	p3.hobby = []string{"篮球", "足球"}
	fmt.Println(p3.name, p3.age, p3.hobby)
}
