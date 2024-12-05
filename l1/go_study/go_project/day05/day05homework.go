package day05

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// 作业（21题）
// 1.Go语言中int占位多少字节
//
// 在32位上占4个字节，在64位上占8个字节
// 2.整型中有符号和无符号是什么意思？
// 无符号整数类型：只能表示非负整数
// 有符号整数类型：表示正数，负数，0
// 3. 整型可以表示最大的范围是多少？ 超出怎么办？
// 超过范围不能正确的表示
// 4. 什么是nil
//
// 5. 十进制是以整型方式存在，其他进制则是以字符串的形式存在？如何实现进制之间的转换
// 6. 简述如下代码的意义
// var v1 int
// var v2 *int
// var v3 = new(int)
// 7. 浮点型为什么有时无法精确表示小数?
// 8. 如何使用第三方包decimal？
// 9. 简述ascii，unicode，utf-8的关系
// 10. 判断：Go语言中的字符串是utf-8编码的字节序列
// 11. 什么是rune?
// 12. 判断:字符串是否可变？
// 13. 列表你记得的字符串的常见功能
// 14. 字符串和字符集合，rune集合如何实现相互转换
// 15. 字符串如何实现高效的拼接？
// 16. 简述数组的存储原理
// 17. 根据要求写代码
// names := [3]string{"Alex","aa","bb"}
// a.请根据索引获取“bb”
// b.请根据索引获取“alex”
// c.请根据索引获取“aa”
// d.请将name数组的最后一个元素修改为“大烧饼"
// 18. 看代码输出结果
// var nestData [3][2]int
// fmt.Println(nestData)
// 19. 请声明一个有3个元素的数组，元素的类型是有两个元素的数组，并在数组中初始化值
// [
// ["alex","aaa"],
// ["eric","admin"],
// ["tony","ppp"],
// ]
// 20. 循环如下数组并使用字符串格式化输出如下内容：
// dataList := {
// ["alex","aaa"],
// ["eric","admin"],
// ["tony","ppp"],
// }
// 最终实现输出:
// 我是Alex,我的密码是aaa.
// 我是eric,我的密码是admin.
// 我是tony,我的密码是ppp.
// 21. 补充代码实现用户登录
// dataList := {
// ["alex","aaa"],
// ["eric","admin"],
// ["tony","ppp"],
// }
func Day05homework() {
	//1.Go语言中int占位多少字节  --在64位系统中占8个字节，32位系统4个字节
	var v1 int = 5
	fmt.Println("size of int:", unsafe.Sizeof(v1)) //size of int: 8
	// 2.整型中有符号和无符号是什么意思？  --uint和int
	// 无符号整数类型：只能表示非负整数
	// 有符号整数类型：表示正数，负数，0
	//var v2 uint = -1 //报错
	var v2 uint = 555
	var v3 int = 343
	fmt.Println(v2, v3)
	// 3. 整型可以表示最大的范围是多少？ 超出怎么办？
	// int64占用8个字节,uint64占用8个字节,超过范围不能正确的表示，使用big包
	// 4. 什么是nil  --nil是空值
	// 5. 十进制是以整型方式存在，其他进制则是以字符串的形式存在？如何实现进制之间的转换
	//在 Go 语言中，十进制数是以整型（int）的形式存在的，你可以直接进行算术运算。而其他的进制数，比如二进制、八进制和十六进制，在 Go 中是以字符串的形式存在的。
	//如果你想要将这些进制的数字转换为十进制，你可以使用 Go 语言的内置函数。对于二进制、八进制和十六进制，你可以使用 strconv.ParseInt 函数来转换为十进制。
	//二进制转换十进制
	binary := "1011"
	binaryInt, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Printf("二进制数%s 转换为十进制%d \n", binary, binaryInt)
	//八进制转十进制
	octal := "123"
	octalInt, _ := strconv.ParseInt(octal, 8, 64)
	fmt.Printf("八进制数%s 转换为十进制%d \n", octal, octalInt)
	//十六进制进制转十进制
	hexdecimal := "1A3"
	hexdecimalInt, _ := strconv.ParseInt(octal, 16, 64)
	fmt.Printf("十六进制数%s 转换为十进制%d \n", hexdecimal, hexdecimalInt)
	//十进制转换为二进制、八进制、十六进制
	decimal := 1111
	binary1 := fmt.Sprintf("%b", decimal)
	fmt.Printf("十进制数%d 转换为二进制%s \n", decimal, binary1)
	octal1 := fmt.Sprintf("%o", decimal)
	fmt.Printf("十进制数%d 转换为八进制%s \n", decimal, octal1)
	hexdecimal1 := fmt.Sprintf("%x", decimal)
	fmt.Printf("十进制数%d 转换为十六进制%s \n", decimal, hexdecimal1)
	// 6. 简述如下代码的意义
	// var v1 int   定义整型 v1
	// var v2 *int  定义指针整型 v2
	// var v3 = new(int)   定义指针类型v3
	// 7. 浮点型为什么有时无法精确表示小数?
	//浮点计算是无法精确表示小数
	// 8. 如何使用第三方包decimal？
	//go语言内部无decimal
	//    第三方包，则需要再本地的go环境中先安装，再使用，第三方包源码地址：https://
	//    go get github.com/shopspring/decimal
	// 9. 简述ascii，unicode，utf-8的关系
	//ASCII‌：ASCII（American Standard Code for Information Interchange）是一种基于拉丁字母的编码系统，主要用于显示现代英语和其他西欧语
	//Unicode是一种国际化的字符编码标准，旨在统一各种语言的编码。它通常使用两个字节表示一个字符，但也支持更大的字符集，包括所有已知语言的字符‌
	//UTF-8‌：UTF-8是一种变长的编码方式，可以表示Unicode中的所有字符。它使用1到4个字节来表示一个字符，英文字符通常使用一个字节，而中文等字符则使用更多的字节‌3。

	// 10. 判断：Go语言中的字符串是utf-8编码的字节序列
	str1 := "hello, 世界" //utf-8编码
	if utf8.ValidString(str1) {
		fmt.Println("str1 is a valid UTF-8 string")
	} else {
		fmt.Println("str1 is not a valid UTF-8 string")
	}
	// 11. 什么是rune?
	//rune表示单个字符（即一个Unicode码点，而字符串string可以由多个字符组成）

	// 12. 判断:字符串是否可变？
	//字符串是不可变的

	// 13. 列表字符串的常见功能
	//字符串常见功能
	//获取长度
	name := "罗曼曼"
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
	//是否以xx开头
	result := strings.HasPrefix(name, "罗")
	fmt.Println(result) //true or false
	//是否以xx结尾
	result1 := strings.HasSuffix(name, "罗")
	fmt.Println(result1) //true or false
	//是否包含
	result2 := strings.Contains(name, "曼")
	fmt.Println(result2) //true or false
	//变大写,小写
	result3 := strings.ToUpper(name)
	fmt.Println(result3)
	result4 := strings.ToLower(name)
	fmt.Println(result4)
	//去两边
	result5 := strings.TrimRight(name, "罗")
	fmt.Println(result5)
	result6 := strings.TrimLeft(name, "曼")
	fmt.Println(result6)
	result7 := strings.Trim(name, "曼")
	fmt.Println(result7)

	//替换
	result8 := strings.Replace(name, "罗", "陈", 1)   //找到罗替换为陈，从左到右找到第一个替换
	result9 := strings.Replace(name, "曼", "陈", 2)   //找到罗替换为陈，从左到右找到前两个替换
	result10 := strings.Replace(name, "曼", "菲", -1) //找到罗替换为陈，替换所有
	fmt.Println(result8, result9, result10)
	//分割
	name1 := "我爱我的家乡河南"
	result11 := strings.Split(name1, "的")
	fmt.Println(result11)
	//拼接
	//不建议
	message := "我爱" + "北京天安门"
	fmt.Println(message)
	//效率高一些
	dataList := []string{"我爱", "北京天安门", "I love you"}
	result12 := strings.Join(dataList, "-")
	fmt.Println(result12)
	//效率更高 (go 1.10之前)
	var buffer bytes.Buffer
	buffer.WriteString("哈哈哈")
	buffer.WriteString("我爱")
	buffer.WriteString("北京天安门")
	data := buffer.String()
	fmt.Println(data)
	//效率更高 (go 1.10之后)
	var builder strings.Builder
	builder.WriteString("哈哈哈")
	builder.WriteString("我爱")
	builder.WriteString("北京天安门")
	data3 := builder.String()
	fmt.Println(data3)
	//string转换为int
	num := "666"
	var data4, _ = strconv.Atoi(num)
	fmt.Println(data4)
	var result20, err1 = strconv.ParseInt(num, 10, 32)
	fmt.Println(result20, err1)

	// 14. 字符串和字符集合，rune集合如何实现相互转换
	//字符串和rune集合
	tempSet2 := []rune(name)
	fmt.Println(tempSet2)
	fmt.Println(tempSet2[0], strconv.FormatInt(int64(tempSet2[0]), 16))
	fmt.Println(tempSet2[1], strconv.FormatInt(int64(tempSet2[1]), 16))
	fmt.Println(tempSet2[2], strconv.FormatInt(int64(tempSet2[2]), 16))
	runeList3 := []rune{32599, 26364, 26364}
	targetString3 := string(runeList3)
	fmt.Println(targetString3)

	// 15. 字符串如何实现高效的拼接？
	dataList1 := []string{"我爱", "北京天安门", "I love you"}
	result33 := strings.Join(dataList1, "-")
	fmt.Println(result33)
	// 16. 简述数组的存储原理
	//数组，定长且元素类型一致的数据集合
	// 17. 根据要求写代码
	// names := [3]string{"Alex","aa","bb"}
	// a.请根据索引获取“bb”
	// b.请根据索引获取“alex”
	// c.请根据索引获取“aa”
	// d.请将name数组的最后一个元素修改为“大烧饼"
	names := [3]string{"Alex", "aa", "bb"}
	fmt.Println(names[2])
	fmt.Println(names[0])
	fmt.Println(names[1])
	names[2] = "大烧饼"
	fmt.Println(names)
	// 18. 看代码输出结果
	var nestData [3][2]int
	fmt.Println(nestData) //[[0 0] [0 0] [0 0]]
	// 19. 请声明一个有3个元素的数组，元素的类型是有两个元素的数组，并在数组中初始化值
	// [
	// ["alex","aaa"],
	// ["eric","admin"],
	// ["tony","ppp"],
	// ]

	nestData1 := [3][2]string{{"alex", "aaa"},
		{"eric", "admin"},
		{"tony", "ppp"}}
	fmt.Println(nestData1)
	// 20. 循环如下数组并使用字符串格式化输出如下内容：
	// dataList := {
	// ["alex","aaa"],
	// ["eric","admin"],
	// ["tony","ppp"],
	// }
	// 最终实现输出:
	// 我是Alex,我的密码是aaa.
	// 我是eric,我的密码是admin.
	// 我是tony,我的密码是ppp.
	nestData2 := [3][2]string{{"alex", "aaa"},
		{"eric", "admin"},
		{"tony", "ppp"}}
	for i := 0; i < len(nestData2); i++ {
		fmt.Printf("我是%s,我的密码是%s. \n", nestData2[i][0], nestData2[i][1])
	}
	// 21. 补充代码实现用户登录
	// dataList := {
	// ["alex","aaa"],
	// ["eric","admin"],
	// ["tony","ppp"],
	// }
	nestData3 := [][]string{{"alex", "aaa"},
		{"eric", "admin"},
		{"tony", "ppp"}}

	var username, pwd string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&username)
	fmt.Print("请输入密码：")
	fmt.Scanln(&pwd)
	if isAuthenticated(nestData3, username, pwd) {
		fmt.Println("登录成功！")
	} else {
		fmt.Println("用户名或密码错误！")
	}
}

func isAuthenticated(users [][]string, username, password string) bool {
	for _, user := range users {
		if user[0] == username && user[1] == password {
			return true
		}
	}
	return false
}
