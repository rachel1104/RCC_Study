package day07

import "fmt"

// 结构体是一个复合类型，用于表示一组数据
// 结构体由一系列属性组成，每个属性都有自己的类型和值
// 定义结构体
type Address struct {
	city, state string
}
type Person1 struct {
	name    string
	age     int
	address Address
}

func Struct1() {
	//先后顺序
	var p1 = Person1{"luomanfei", 19, Address{"北京", "中国"}}
	fmt.Println(p1.name, p1.age, p1.address.city)
	//关键字
	var p2 = Person1{name: "luomanfei", age: 19, address: Address{"北京", "中国"}}
	fmt.Println(p2.name, p2.age, p2.address)
	//先声明再赋值
	var p3 Person1
	p3.name = "luomanfei"
	p3.age = 19
	p3.address = Address{"北京", "中国"}
	fmt.Println(p3.name, p3.age, p3.address)
}