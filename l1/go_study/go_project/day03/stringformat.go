package day03

import "fmt"

func FormatString() {
	var name, address, action string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入位置：")
	fmt.Scanln(&address)
	fmt.Print("请输入行为：")
	fmt.Scanln(&action)
	result := fmt.Sprintf("我叫%s,在%s正在干%s", name, address, action)
	fmt.Println(result)

}
