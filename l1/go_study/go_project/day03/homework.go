package day03

import "fmt"

// 1.猜年龄游戏，要求：允许用户最多尝试三次，3次都没猜对的话，就直接退出，如果猜对了，打印恭喜信息并退出
// 2.实现用户登录系统，并且要求支持连续3次输错之后直接退出，并且每次输错时显示剩余次数（提示：使用字符串格式化输出）
func Day03Homework() {
	////1.猜年龄游戏，要求：允许用户最多尝试三次，3次都没猜对的话，就直接退出，如果猜对了，打印恭喜信息并退出
	//for i := 1; i <= 3; i++ {
	//	var age, guessAge int
	//	age = 20
	//	fmt.Print("请输入年龄：")
	//	fmt.Scan(&guessAge)
	//	if guessAge == age {
	//		fmt.Println("恭喜你猜对了")
	//		break
	//	} else {
	//		fmt.Printf("您第%d次没猜对", i)
	//		fmt.Println("")
	//	}
	//}
	//2.实现用户登录系统，并且要求支持连续3次输错之后直接退出，并且每次输错时显示剩余次数（提示：使用字符串格式化输出）
	for i := 1; i <= 3; i++ {
		var user, password, guessUser, guessPassword string
		user = "admin"
		password = "admin123"
		fmt.Print("请输入用户名：")
		fmt.Scan(&guessUser)
		fmt.Print("请输入密码：")
		fmt.Scan(&guessPassword)
		if guessUser == user && guessPassword == password {
			fmt.Println("登录成功！")
			break
		} else {
			if i == 3 {
				fmt.Printf("用户已锁定，请联系管理员!")
				break
			}
			fmt.Printf("您输入的用户名和密码错误，还剩%d次机会！", 3-i)
			fmt.Println("")
		}
	}
}
