package day07

import "fmt"

/*
函数

	func 函数名（参数） 返回值 {
		函数体
	}

函数命名时需注意：函数名只能是字母，数字，下划线组合且数字不能开头（驼峰式命名）
函数可以做参数
*/
func add100(arg int) (int, bool) {
	return arg + 100, true
}

//	func Function1(data int, exec func(int) (int, bool)) int {
//		data, flag := exec(data)
//		if flag {
//			return data
//		} else {
//			return 9999
//		}
//	}
type f1 func(arg int) (int, bool)

func Function1(data int, exec f1) int {
	data, flag := exec(data)
	if flag {
		return data
	} else {
		return 9999
	}
}

// 变长参数,变长参数只能放在最后，且只能有一个
func Do(num ...int) int {
	fmt.Println(num)
	sum := 0
	for _, num := range num {
		sum += num
	}
	return sum
}
func main() {
	result := Function1(100, add100)
	fmt.Println(result)
	r1 := Do(1, 2, 3, 3)
	r2 := Do(1, 2, 3)
	fmt.Println(r1, r2)
}
