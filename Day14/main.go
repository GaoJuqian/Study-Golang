package main

import "fmt"

// 学生信息管理系统

/*
	1.添加学生信息
	2.编辑学生信息
	3.查看所有学生信息
*/

func main() {
	// 初始化学生列表
	stuList := NewStudentList()
	for {
		fmt.Println(`学生信息管理系统；1.添加学生信息，2.编辑学生信息，3.查看所有学生信息`)
		var input int
		fmt.Print("请输入操作：")
		fmt.Scanf("%d\n", &input)

		var (
			id        int
			name, sex string
		)

		switch input {
		case 1:
			// 添加学生
			fmt.Print("请输入id：")
			fmt.Scanf("%d\n", &id)
			fmt.Print("请输入NAME：")
			fmt.Scanf("%s\n", &name)
			fmt.Print("请输入SEX：")
			fmt.Scanf("%s\n", &sex)
			// 添加学生信息
			stuList.AddStudent(id, name, sex)
		case 2:
			// id > 修改学生信息
			fmt.Print("请输入修改的id：")
			fmt.Scanf("%d\n", &id)
			fmt.Print("请输入NAME：")
			fmt.Scanf("%s\n", &name)
			fmt.Print("请输入SEX：")
			fmt.Scanf("%s\n", &sex)
			stuList.EditStudent(id, name, sex)
		case 3:
			// 查询所有学生信息
			stuList.FindAll()
		}
		fmt.Println("操作成功")
		//fmt.Printf("%#v\n %v\n %v\n", stuList, len(stuList.List), cap(stuList.List))
		//for _, v := range stuList.List {
		//	fmt.Printf("%#v\n", v)
		//}

	}
}
