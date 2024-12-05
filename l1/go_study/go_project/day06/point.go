package day06

import (
	"fmt"
	"unsafe"
)

func Point() {
	//指针
	v1 := "luomanfei"
	v2 := &v1
	fmt.Println(v1, v2, *v2) //luomanfei 0xc0000240a0 luomanfei

	v1 = "alex"
	fmt.Println(v1, v2, *v2) //alex 0xc0000240a0 alex

	//指针的指针
	v3 := "luomanfei"
	var p1 *string = &v3
	var p2 **string = &p1
	var p3 ***string = &p2
	fmt.Println(v3, &v3)
	fmt.Println(p1, &p1)
	fmt.Println(p2, &p2)
	fmt.Println(p3, &p3)
	//指针的计算
	dataList := [3]int8{11, 22, 33}
	//获取数组第一个元素的地址（指针）
	var firstDataPtr *int8 = &dataList[0]
	//转换成Point类型
	ptr := unsafe.Pointer(firstDataPtr)
	//转换成uintptr类型，然后进行内存地址的计算
	targetAddress := uintptr(ptr) + 1
	//根据新地址，重新转换成Point类型
	newPtr := unsafe.Pointer(targetAddress)
	//Point对象转换为int8指针类型
	value := (*int8)(newPtr)
	//根据指针获取取值
	fmt.Println("最终结果为：", *value) //最终结果为：22
}
