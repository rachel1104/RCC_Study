package day02

import "fmt"

/*
1.提示用户输入麻花腾，判断用户输入的对不对，如果对，提示真聪明，如果不对提示，你的智商待提高
2.提示用户输入两个数字，并计算两个数的和并输出
3.提示用户输入姓名，位置，行为三个值，然后做字符串的拼接并输出，例如：xx在xx做xx
4.设定一个理想数字比如：66，让用户输入数字，如果比66大，则显示猜测的结果大了，如果比66小，则显示猜测的结果小了，
如果等于66，显示猜测结果正确
5.写程序，输出成绩等级，成绩有ABCDE5个等级，与分数的相应关系如下
A：90-100
B：80-89
C：60-79
D：40-59
E：0-39
要求用户输入0-100的数字后，正确输入他的对应成绩等级
*/

/*
声明+赋值
var name string = "test"
var name = "test"
name := "test"   推荐
*/
/*
多个变量
var (
	name = "rachel"
	age = "18"
	hobby = "大保健"
	salary =100000
	gender string
	length int
	sb bool
)
*/
/*
作用域：
如果我们定义了大括号，那么大括号中定义的变量
1.不能被上级使用
2.可以在同级使用
3.可以在子级使用
全局变量和局部变量
全局变量:未写在函数中的变量成为全局变量，不可以使用v1:=xx方式简化，
可以基于饮食分解方式声明多个变量
*/
func Grade() {
	var name string
	fmt.Println("请输入麻花藤，")
	fmt.Scanf("%s", &name)
	if name == "麻花藤" {
		fmt.Println("真聪明！")
	} else {
		fmt.Println("你的智商待提高！")
	}

}
func InputNumber() {
	var num1, num2 int
	fmt.Println("请输入第一个数字")
	fmt.Scanf("%d", &num1)
	fmt.Println("请输入第二个数字")
	fmt.Scanf("%d", &num2)
	fmt.Println(num1 + num2)
}
func Userinfo() {
	var username, location, doing string
	fmt.Println("请输入用户名：")
	fmt.Scanf("%s", &username)
	fmt.Println("请输入位置：")
	fmt.Scanf("%s", &location)
	fmt.Println("请输入行为：")
	fmt.Scanf("%s", &doing)
	fmt.Println(username + "在" + location + doing)
}
func Guess() {
	var num int
	fmt.Println("请输入数字：")
	fmt.Scanf("%d", &num)

	if num > 66 {
		fmt.Println("猜测的结果大了")
	} else if num == 66 {
		fmt.Println("数字你猜对了")
	} else {
		fmt.Println("猜测的结果小了")
	}
}

func StudentGuess() {
	var grade int
	fmt.Println("请输入成绩：")
	fmt.Scanf("%d", &grade)
	if grade >= 90 && grade <= 100 {
		fmt.Println("A")
	} else if grade >= 80 && grade <= 89 {
		fmt.Println("B")
	} else if grade >= 60 && grade <= 79 {
		fmt.Println("C")
	} else if grade >= 40 && grade <= 59 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

/*
name,nickname的内存不一样,
使用int，string，bool这三种数据类型时，如果遇到变量的赋值则会copy一份
*/
func Mod() {
	name := "222aaa"
	nickname := "www"
	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)
}
