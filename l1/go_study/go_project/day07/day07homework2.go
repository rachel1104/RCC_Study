package day07

import "fmt"

// 2.创建班级，并展示班级信息
// 创建学校和班级的结构体，默认创建一个学校对象
// 根据用户输入去创建班级（班级名称和人数），然后将创建的班级信息添加到一个切片中
type School1 struct {
	band string
	city string
}

type Class1 struct {
	title  string
	count  int
	school School1
}

func Day07homework2() {
	//1.创建学校
	sch := School1{"aaa", "beijing"}
	//2.创建学校
	c1 := Class1{"python", 80, sch}
	c2 := Class1{"c#", 70, sch}
	c1.school.band = "big boy"
	fmt.Println(c1, c2)
}
