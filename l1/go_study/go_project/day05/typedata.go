package day05

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func DataType() {
	////整型转换为字符串类型
	//v1 := 19
	//result := strconv.Itoa(v1)
	//fmt.Println(result, reflect.TypeOf(result))
	//var v2 int8 = 17
	//data := strconv.Itoa(int(v2))
	//fmt.Println(data, reflect.TypeOf(result))
	////字符串转换为整型
	//v3 := "666"
	//result3, err := strconv.Atoi(v3)
	//fmt.Println(result3, err)
	//if err == nil {
	//	fmt.Println("转换成功")
	//} else {
	//	fmt.Println("转换失败")
	//}

	////整型转换为二进制
	//v1 := 5
	//r1 := strconv.FormatInt(int64(v1), 2)
	//r2 := strconv.FormatInt(int64(v1), 8)
	//r3 := strconv.FormatInt(int64(v1), 16)
	//fmt.Println(r1, reflect.TypeOf(r1))
	//fmt.Println(r2, reflect.TypeOf(r2))
	//fmt.Println(r3, reflect.TypeOf(r3))

	//data := "1001000101"
	//result, err := strconv.ParseInt(data, 2, 16)
	//fmt.Println(result, err)

	////常见数学运算
	//fmt.Println(math.Abs(-19))               //取绝对值
	//fmt.Println(math.Floor(3.14))            //向下取整
	//fmt.Println(math.Ceil(3.14))             //向上取整
	//fmt.Println(math.Round(3.378*100) / 100) //就近取整
	//fmt.Println(math.Mod(11, 3))             //取余数
	//fmt.Println(math.Pow(2, 5))              //计算次方 2的5次方
	//fmt.Println(math.Pow10(2))               //计算10次方，2的10次方
	//fmt.Println(math.Max(1, 2))              //两个值，取大
	//fmt.Println(math.Min(1, 2))              //两个值，取小

	//第一步：创建一个超大整型的一个对象
	//var v1 big.Int //不能直接写值，一般用不到，直接赋值是才使用
	//fmt.Println(v1)
	//var v2 *big.Int
	//fmt.Println(v2)
	//v3 := new(big.Int)
	//fmt.Println(v3)
	////第二步：在超大整型对象中写入值
	////v1.SetInt64(1990)
	//v1.SetString("323234234234234234234123", 10) //10代表10进制，2代表二进制
	//print(v1)
	//n1 := new(big.Int)
	//n1.SetInt64(89)
	//n2 := new(big.Int)
	//n2.SetInt64(99)
	//指针
	//n1 := big.NewInt(89)
	//n2 := big.NewInt(99)
	////minder := new(big.Int)
	//result := new(big.Int)
	//result.Add(n1, n2)
	//
	//result.Sub(n1, n2)
	//result.Mul(n1, n2)
	//result.Div(n1, n2)
	//result.DivMod(n1, n2, minder)
	//fmt.Println(result.String())

	//float32，float64,处理小数时非精确的表示
	//var a1 float32
	//a1 = 3.14
	//a2 := 99.9 //默认float64，需转换
	//a3 := float64(a1) + a2
	//fmt.Println(a1, a2, a3)
	////decimal
	//var price = decimal.NewFromFloat(3.4626)
	//var data1 = price.Round(1)    //保留小数后1位
	//var data2 = price.Truncate(1) //保留小数后1位
	//fmt.Println(data1, data2)     //输出:3.5， 3.4
	//字符串转布尔
	//result, err := strconv.ParseBool("t")
	//fmt.Println(result, err)
	//布尔转字符串
	//result := strconv.FormatBool(true)
	//fmt.Println(result)
	//string
	var name string = "罗曼曼"
	//罗 =>
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))
	//曼
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))
	//曼
	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	fmt.Println(name)
	//获取字符串的长度:9，字节长度
	fmt.Println(len(name))
	//字符串转换为一个字节合集
	byteSet := []byte(name) //[231 189 151 230 155 188 230 155 188]
	fmt.Println(byteSet)
	byteList := []byte{231, 189, 151, 230, 155, 188, 230, 155, 188}
	targetString := string(byteList)
	fmt.Println(targetString)
	//5. rune的集合
	tempSet := []rune(name)
	fmt.Println(tempSet) //[32599 26364 26364]
	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))
	//6. rune集合 转换为字符串
	runeList := []rune{32599, 26364, 26364}
	targetName := string(runeList)
	fmt.Println(targetName)
	//7.长度的处理
	runeLength := utf8.RuneCountInString(name)
	fmt.Println(runeLength)
	//字符串常见功能
	//获取长度
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

	//int转string
	var result21 = strconv.Itoa(888)
	fmt.Println(result21)
	//字符串和字节集合
	var name5 string = "罗曼曼"
	byteSet1 := []byte(name5)
	//字符串转换为一个字节集合
	fmt.Println(byteSet1)
	//字节的集合转换为字符串
	byteList1 := []byte{231, 189, 151, 230, 155, 188, 230, 155, 188}
	targetString2 := string(byteList1)
	fmt.Println(targetString2)
	//字符串和rune集合
	tempSet2 := []rune(name5)
	fmt.Println(tempSet2)
	fmt.Println(tempSet2[0], strconv.FormatInt(int64(tempSet2[0]), 16))
	fmt.Println(tempSet2[1], strconv.FormatInt(int64(tempSet2[1]), 16))
	fmt.Println(tempSet2[2], strconv.FormatInt(int64(tempSet2[2]), 16))
	runeList3 := []rune{32599, 26364, 26364}
	targetString3 := string(runeList3)
	fmt.Println(targetString3)
	//string 和字符
	v1 := string(65)
	fmt.Println(v1) //A
	v2, size := utf8.DecodeRuneInString("A")
	fmt.Println(v2, size)
	//索引切片和循环
	//1.索引获取字节
	var name6 string = "布丁小哲"
	c1 := name6[0] //(最大11，没个文字包含3)
	fmt.Println(c1)
	//2.切片获取字节区间
	c2 := name6[0:3]
	fmt.Println(c2)
	//3.手动循环获取所有字节
	for i := 0; i < len(name6); i++ {
		fmt.Println(i, name6[i])
	}
	//4.for range循环获取所有字符
	for index, item := range name6 {
		fmt.Println(index, item, string(item))
	}
	//5.转换成rune集合
	dataList5 := []rune(name6)
	fmt.Println(dataList5[0], string(dataList5[0]))

	//数组，定长且元素类型一致的数据集合
	//先声明再赋值
	var numbers [3]int
	numbers[0] = 999
	numbers[1] = 666
	numbers[2] = 333
	//声明+赋值
	var names = [2]string{"曼曼", "布丁"}
	//声明+赋值+指定位置
	var ages = [3]int{0: 87, 1: 73, 2: 99}
	//省略个数
	var names1 = [...]string{"曼曼", "布丁"}
	var ages1 = [...]int{0: 87, 1: 73, 2: 99}
	//声明 指针类型的数组，不会开辟内存初始化数组的值
	var numbers3 *[3]int
	//声明数组并初始化
	numbers4 := new([3]int)
	fmt.Println(names, ages, names1, ages1, numbers3, numbers4)

	nums := [3]int8{11, 22, 33}
	fmt.Println("数组的内存地址：%p \n", &nums)
	fmt.Println("数组第一个元素的内存地址：%p \n", &nums[0])
	fmt.Println("数组第二个元素的内存地址：%p \n", &nums[1])
	fmt.Println("数组第三个元素的内存地址：%p \n", &nums[2])
	//可变和copy
	//数组的元素可以被更改
	names2 := [2]string{"罗曼", "布丁"}
	names2[1] = "罗哲"
	//copy
	names3 := names2
	fmt.Println(names3)
	//长度，索引切片和循环
	fmt.Println(len(names2))
	data5 := names2[0]
	fmt.Println(data5)
	names2[0] = "eric"
	fmt.Println(names2)
	data6 := names2[0:2] //获取0-2的下标的数据
	fmt.Println(data6)
	//循环
	nums8 := [3]int32{11, 22, 33}
	for i := 0; i < len(nums8); i++ {
		fmt.Println(i, nums8[i])
	}
	for key, item := range nums8 {
		fmt.Println(key, item)
	}
	//数组嵌套
	var nestData [2][3]int
	nestData[0] = [3]int{11, 22, 33}
	nestData[1][1] = 666
	fmt.Println(nestData)

}
