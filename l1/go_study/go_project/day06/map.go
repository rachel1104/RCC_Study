package day06

import "fmt"

// 字典（dict）或映射(map)
// MAP特点：
// 键不能重复，键必须可hash（可hash的有int/bool/float/string/array），无序
func Map() {
	////声明
	//userInfo1 := map[string]string{}
	//声明和初始化
	//userInfo := map[string]string{"name": "罗曼粉", "age": "18"}
	////name1 = userInfo["name"] //罗曼粉
	//userInfo["age"] = "20"
	//userInfo["email"] = "waa@live.com"
	//fmt.Println(userInfo)
	//fmt.Println(len(userInfo))

	//嵌套
	//data1 := make(map[string]string, 10)
	//data2 := make(map[string]int, 10)
	//data3 := make(map[string][2]int, 10)
	//增删改查
	data := map[string]string{"name": "罗曼粉", "age": "18"}
	data["email"] = "waa@live.com" //添加
	data["age"] = "16"             //修改
	//delete(data, "age")            //删除
	for k, v := range data {
		fmt.Println(k, v)
	}

}
