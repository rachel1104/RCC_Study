package day03

import "fmt"

/*
switch case语句，条件判断
for 循环语句
goto语法，不太建议使用
字符串格式化
*/
func Yuju() {
	////表达式
	//switch 1 + 1 {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2")
	//case 3:
	//	fmt.Println("3")
	//default:
	//	fmt.Println("都不满足")
	//}

	//变量
	//age := 18
	//var age int
	//fmt.Scanln(&age)
	//switch age {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println("2")
	//case 3:
	//	fmt.Println("3")
	//case 18:
	//	fmt.Println("18")
	//default:
	//	fmt.Println("都不满足")
	//}
	//注意事项：字符类型要保持一致

	//for 循环 死循环
	//for {
	//	fmt.Println("红鲤鱼与绿鲤鱼与驴")
	//	time.Sleep(time.Second * 1)
	//}

	//for 有条件的循环
	//number := 1
	//for number < 5 {
	//	fmt.Println("红鲤鱼与绿鲤鱼与驴")
	//	number = number + 1
	//}

	for i := 1; i <= 5; i++ {
		fmt.Println("红鲤鱼与绿鲤鱼与驴")
	}
}
