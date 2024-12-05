package day06

import "fmt"

/*
1、简述切片和数组的区别
在Go语言中，数组和切片都可以用来存储一系列元素，但它们之间有一些关键的区别：
数组（Array）是固定长度的，而切片（Slice）是灵活的、动态的。
数组是值类型，而切片是引用类型。
数组的长度是固定的，不能改变，而切片的长度是其底层数组长度和其位置的长度之和。
切片可以通过append()函数进行元素添加，当切片容量不足时，会自动扩展。

2.简述new和make的区别
new:用于分配内存并返回指向该内存的指针
make:用于初始化和分配特定类型的内存，只能用于slice,map和channel
3.看代码写结果
v1 := make([]int,2,5)
fmt.Println(v1[0],len(v1),cap(v1))
4.看代码写结果
*/
func Slicehomework() {
	//v1 := make([]int, 2, 5)
	//fmt.Println(v1[0], len(v1), cap(v1)) //0,2,5

	//v1 := make([]int, 2, 5)
	//v2 := append(v1, 123)
	//fmt.Println(len(v1), cap(v1)) // 2,5
	//fmt.Println(len(v2), cap(v2)) // 3,5

	//v1 := make([]int, 2, 5)
	//v2 := append(v1, 123)
	//v1[0] = 999
	//fmt.Println(v1) // [999 0]
	//fmt.Println(v2) //[999 0 123]

	//v1 := make([]int, 2, 2)
	//v2 := append(v1, 123)
	//v1[0] = 999
	//fmt.Println(v1) // [999 0]
	//fmt.Println(v2) //[0 0 123]

	//v1 := make([]int, 2, 2)
	//v2 := v1[0:2]
	//v1[0] = 111
	//fmt.Println(v1) // [111 0]
	//fmt.Println(v2) //[111 0]

	//v1 := [][]int{[]int{11, 22, 33, 44}, []int{44, 55}}
	//v1[0][2] = 999
	//fmt.Println(v1) // [[11 22 999 44] [44 55]]

	//v1 := [][]int{[]int{11, 22, 33, 44}, []int{44, 55}}
	//v2 := v1[0]
	//v2[2] = 60
	//fmt.Println(v1) // [[11 22 60 44] [44 55]]

	//v1 := [][]int{[]int{11, 22, 33, 44}, []int{44, 55}}
	//v2 := append(v1[0], 99) //扩容新地址
	//v2[2] = 69
	//fmt.Println(v1) // [[11 22 33 44] [44 55]]
	//fmt.Println(v2) // [[11 22 33 44] [44 55]]

	v1 := [][]int{make([]int, 2, 5), make([]int, 2, 3)}
	v2 := append(v1[0], 99) //扩容新地址
	fmt.Println(v1)         // [[0 0] [0 0]]
	fmt.Println(v2)         // [0 0 99]
	v2[0] = 69
	fmt.Println(v1) // [[69 0] [0 0]]
	fmt.Println(v2) // [69 0 99]
}
