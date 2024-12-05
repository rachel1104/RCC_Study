package day06

import "fmt"

func Slice() {
	////创建切片
	//var nums []int
	//var users = make([]int, 2, 5) //make只能用于切片，字典，channel
	//var data = []int{11, 22, 33}
	//
	////切片的指针类型
	//var v1 = new([]int)
	//var v2 = *[]int   //指针类型（nil）
	//
	////自动扩容
	//v1 := make([]int,1,3)  //cap=3,len=1,array=0xc000122028
	//fmt.Println(len(v1),cap(v1))
	//v2 := append(v1,99)
	//fmt.Println(v1)  //[0]
	//fmt.Println(v2)  //[0,99]
	////其他
	//v1 := make([]int, 3)
	//切片常见操作
	//1.长度和容量
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(v1), cap(v1))
	////索引
	//v1 := []string{"alex", "李杰", "老男孩"}
	//v1[0]
	//v1[1]
	//v1[2]
	//v2 := make([]int, 2, 5)
	//v2[0]
	//v2[1]
	//v2[2] //报错
	//v2[0] = 999
	//切片
	v3 := []int{1, 2, 3, 4, 5}
	v4 := v3[1:3]
	v5 := v3[1:]
	v6 := v3[:3]
	fmt.Println(v3, v4, v5, v6)
	//追加
	v7 := append(v3, 44)
	v8 := append(v3, 44, 55, 66, 77)
	v9 := append(v3, []int{100, 200, 300}...)
	fmt.Println(v7, v8, v9)
	//删除  --一般不使用
	deleteIndex := 2
	result := append(v3[:deleteIndex], v3[deleteIndex+1:]...)
	fmt.Println(result)
	//插入
	insertIndex := 3 //在索引位置插入99
	//result1 := make([]int, 0, len(v1)+1)
	//result1 = append(result1, v3[:insertIndex]...)
	//result1 = append(result1, 99)
	//result1 = append(result1, v3[insertIndex:]...)
	//fmt.Println(result1)
	result1 := append(v3[:insertIndex], 99)
	result1 = append(result1, v3[insertIndex:]...)
	//循环
	a1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}
	for index, value := range a1 {
		fmt.Println(index, value)
	}
	//变量赋值
	b1 := 1
	b2 := b1
	fmt.Println(&b1, &b2)
}
