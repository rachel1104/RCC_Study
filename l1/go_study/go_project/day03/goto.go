package day03

import "fmt"

func Sgoto() {
	var name string
	fmt.Println("请输入用户名：")
	fmt.Scanln(&name)

	if name == "wwwww" {
		goto SVIP
	} else if name == "bbb" {
		goto VIP
	}

	fmt.Println("预约...")
VIP:
	fmt.Println("等号...")
SVIP:
	fmt.Println("进入...")
}
