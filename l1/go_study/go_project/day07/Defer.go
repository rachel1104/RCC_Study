package day07

import "fmt"

// defer用于在一个函数执行完成后自动触发的语句，一般用于结束操作之后释放资源
// 多个defer时，按照倒序线上
func do() int {
	fmt.Println("风吹")
	defer fmt.Println("函数执行完毕了")
	fmt.Println("aaa")
	return 666
}

// 自执行函数结构体工厂
func main() {
	ret := do()
	fmt.Println(ret)
}
