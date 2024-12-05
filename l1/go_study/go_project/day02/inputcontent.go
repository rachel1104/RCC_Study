package day02

import "fmt"

/*
输入
fmt.Scan: 要求输入两个值，必须输入两个，否则会一直等待
fmt.Scanln: 等待回车，回车后算输入完成
fmt.Scanf: 支持以模板的方式输入
*/
/*
示例1
*/
func InputContent() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	println(name)
}

/*
示例2
*/
func InputContent2() {
	var name string
	var age int
	fmt.Println("Enter your name and age: ")
	_, err := fmt.Scan(&name, &age)
	if err != nil {
		fmt.Println(name, age)
	} else {
		fmt.Println("用户输入数据错误", err)
	}
	//fmt.Println(count, err)
	//fmt.Println(name, age)
}
