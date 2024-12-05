package day07

import "fmt"

/*
返回函数
*/
func exec(num1 int, num2 int) string {
	fmt.Println("执行函数")
	return "成功"
}
func getFunction() func(int, int) string {
	//返回函数
	return exec
}
func main() {
	function := getFunction()
	result := function(111, 222)
	fmt.Println(result)
	//匿名函数
	v1 := func(n1 int, n2 int) int {
		return 123
	}
	data := v1(11, 22)
	fmt.Println(data)
}
