package day02

import "fmt"

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

/*
常量：const ,不能被修改,常量必须赋值
iota : 计数器
*/
func Changliang() {
	const age int = 98
	const salary = 100000
	const (
		v1 = iota
		v2
		v3
		v4
		v5
	)
	fmt.Println(v1, v2, v3, v4, v5, salary)
}

/*
示例
常量：const ,不能被修改,常量必须赋值
iota : 计数器+1
*/
func Changliang1() {
	const (
		v1 = iota + 1
		v2
		v3
		v4
		v5
	)
	fmt.Println(v1, v2, v3, v4, v5)
}

/*
示例
常量：const ,不能被修改,常量必须赋值
iota : 计数器+1  获取2,4,5,6，7
*/
func Changliang2() {
	const (
		v1 = iota + 2
		_
		v2
		v3
		v4
		v5
	)
	fmt.Println(v1, v2, v3, v4, v5)
}
