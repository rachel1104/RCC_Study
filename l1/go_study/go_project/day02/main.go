package day02

import "fmt"

func Gess() {
	fmt.Println("aaa")
	fmt.Println("bbb")
	//整型
	fmt.Println(666)
	fmt.Println(6 + 9)
	fmt.Println(6 * 9)
	fmt.Println(6 / 9)
	fmt.Println(6 % 9)
	//字符串
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("aaa" + "bbb")
	fmt.Println("10" + "2")
	//布尔类型,真假
	fmt.Println(1 > 2)
	fmt.Println(3 > 2)
	if 2 > 1 {
		fmt.Println("真")
	} else {
		fmt.Println("假")
	}
	//变量,声明+赋值
	var sd string = "laonanhai"
	fmt.Println(sd)
	var age int = 73
	fmt.Println(age)
	var flag1 bool = true
	fmt.Println(flag1)
	//先声明再赋值
	var aa string
	aa = "afwradf"
	fmt.Println(aa)
	//存储数据
	var v1 int = 99
	var v2 int = 33
	var v3 int = v1 + v2
	var v4 int = v1 * v2
	var v5 int = v1 + v2 + v3
	fmt.Println(v1, v2, v3, v4, v5)

	//存储用户输入的值
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	if name == "alex" {
		fmt.Println("用户名输入正确")
	} else {
		fmt.Println("用户输入失败")
	}
	//变量名要求
	//变量名必须只包含：字母，数字，下划线, 数字不能开头，不能使用go语音内置的关键字
}
