package day07

import "fmt"

// 1.创建学校结构体，并添加到切片中
// 创建一个学校的结构体，根据用户输入去创建学校（学校包含品牌和城市），然后将创建的学校信息动态添加到一个切片中
type School struct {
	band string
	city string
}

func Day07homework1() {
	var schoolList []School
	for {
		//用户输入品牌和城市
		var band, city string
		fmt.Printf("请输入品牌：")
		fmt.Scanf("%s", &band)
		fmt.Printf("请输入城市：")
		fmt.Scanf("%s", &city)
		if band == "1" {
			break
		}
		//创建一个结构体对象
		sch := School{band: band, city: city}
		schoolList = append(schoolList, sch)
	}
	fmt.Println(schoolList)
}
