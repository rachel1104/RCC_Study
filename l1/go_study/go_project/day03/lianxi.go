package day03

import "fmt"

/*
1.猜数字，设定一个理想数字比如：66，一直提示让用户输入数字，如果比66大，则显示猜测的结果大了，如果比66小，
则显示猜测的结果小了，只有输入等于66，显示猜测结果正确，然后退出循环
2.使用循环输出1-100的所有整数
3.使用循环输出1,2,3,4,5,6,8,9，10 ，10以内除7以外的整数
4.输出1-100内的所有奇数
5.输出1-100内的所有偶数
6.求1-100的所有整数的和
7.输出10-1所有整数
*/
func Lianxi() {
	////1.猜数字，设定一个理想数字比如：66，一直提示让用户输入数字，如果比66大，则显示猜测的结果大了，如果比66小，
	//data := 66
	//flag := true
	//for flag {
	//	var userInputNumber int
	//	fmt.Println("请输入数字：")
	//	fmt.Scan(&userInputNumber)
	//	if userInputNumber > data {
	//		fmt.Println("太大了")
	//	} else if userInputNumber == data {
	//		fmt.Println("猜对了")
	//		flag = false
	//	} else {
	//		fmt.Println("太小了")
	//	}
	//}
	//
	////2.使用循环输出1-100的所有整数
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	////3.使用循环输出1,2,3,4,5,6,8,9，10 ，10以内除7以外的整数
	//for i := 1; i <= 10; i++ {
	//	if i != 7 {
	//		fmt.Println(i)
	//	}
	//
	//}

	//4.输出1-100内的所有奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println("奇数为：", i)
		}
	}

	//5.输出1-100内的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("偶数为：", i)
		}
	}

	//6.求1-100的所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和是：", sum)

	//7.输出10-1所有整数
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
