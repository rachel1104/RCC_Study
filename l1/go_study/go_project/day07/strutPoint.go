package day07

import "fmt"

type Person4 struct {
	name string "姓名" //”姓名“为标签
	age  int    "年龄"
}

func StrutPoint() {
	p1 := Person4{"test", 18}
	fmt.Println(p1)
	p2 := &Person4{"test", 18}
	fmt.Println(p2.name, p2.age)
}
