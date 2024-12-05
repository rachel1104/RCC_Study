package day03

import "fmt"

func ContinueYuju() {
	//for {
	//	fmt.Println("Alex今天不在家")
	//	continue
	//	fmt.1Println("让谁一起")
	//}
	//for i := 1; i <= 10; i++ {
	//	if i == 7 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//
	////for循环嵌套
	//for i := 1; i <= 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Println(i, j)
	//	}
	//}

	////break
	//for {
	//	fmt.Println("loop")
	//	break
	//	fmt.Println("aaa")
	//}
	////break
	//for i := 1; i <= 10; i++ {
	//	for j := 1; j <= i; j++ {
	//		if j == 7 {
	//			break
	//		}
	//		fmt.Println(i, j)
	//	}
	//}

	//对for打标签，然后通过break和continue就可以实现多层循环的跳出和终止
	//f1:
	//	for i := 1; i <= 10; i++ {
	//		for j := 1; j <= i; j++ {
	//			if j == 7 {
	//				continue f1
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}
	//}

	//对for打标签，然后通过break和continue就可以实现多层循环的跳出和终止
f2:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 7 {
				break f2
			}
			fmt.Println(i, j)
		}
	}
}
