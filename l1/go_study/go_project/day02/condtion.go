package day02

import "fmt"

/*
条件语句

	if true {
		pass
	} else {

		pass
	}
*/

/*
多条件语句

	if len<1 {
		pass
	} else if len<6 {
		pass
	} else {
		pass
	}
*/

/*
嵌套语句

	if len<1 {
		if true {
		pass
		} else {
		}
	} else if len<6 {
		pass
	} else {
		pass
	}
*/
func Cond() {
	var number int
	fmt.Println("请输入数字：")
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("您输入的是偶数")
	} else {
		fmt.Println("您输入的是奇数")
	}

}
