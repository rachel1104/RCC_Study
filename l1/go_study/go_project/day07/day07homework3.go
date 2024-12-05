package day07

import "fmt"

// 2.创建班级，并展示班级信息
// 创建学校和班级的结构体，默认创建一个学校对象
// 根据用户输入去创建班级（班级名称和人数），然后将创建的班级信息添加到一个切片中
type School2 struct {
	band string
	city string
}

type Class2 struct {
	title  string
	count  int
	school *School2
}

func Day07homework3() {
	//1.创建学校
	sch2 := &School2{"aaa", "beijing"}
	var classList []Class2
	//2.循环创建班级
	for {
		var cls Class2
		fmt.Printf("请输入班级：")
		fmt.Scanf("%s", &cls.title)
		if cls.title == "1" {
			break
		}
		fmt.Printf("请输入人数：")
		fmt.Scanf("%d", &cls.count)
		//1.创建班级对象
		cls.school = sch2
		//加入到班级的列表
		classList = append(classList, cls)
	}
	fmt.Println(classList)
	for _, item := range classList {
		message := fmt.Sprintf("%s%s校区，%s班级有%d个学生。", item.school.band, item.school.city, item.title, item.count)
		fmt.Println(message)
	}

}
