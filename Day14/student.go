package main

import "fmt"

// 全部学生
type StudentList struct {
	List []*Student
}

// 学生信息
type Student struct {
	Id        int
	Name, Sex string
}

// 学生构造函数
func NewStudent(id int, name, sex string) *Student {
	return &Student{
		Id:   id,
		Name: name,
		Sex:  sex,
	}
}

func NewStudentList() *StudentList {
	return &StudentList{
		List: make([]*Student, 0, 50),
	}
}

// 添加学生信息
func (stuList *StudentList) AddStudent(id int, name, sex string) {
	stu := NewStudent(id, name, sex)
	//fmt.Printf("%#v\n %v\n %v\n", stuList, len(stuList.List), cap(stuList.List))
	stuList.JoinStuList(stu)
}

// 将单个学生存入列表
func (stuList *StudentList) JoinStuList(stu *Student) {
	stuList.List = append(stuList.List, stu)
	//fmt.Printf("%#v\n %v\n %v\n", stuList, len(stuList.List), cap(stuList.List))
}

// 编辑学生信息
func (stuList *StudentList) EditStudent(id int, name, sex string) {
	//fmt.Printf("%#v\n %v\n %v\n", stuList, len(stuList.List), cap(stuList.List))
	for _, v := range stuList.List {
		if id == v.Id {
			v.Name = name
			v.Sex = sex
			fmt.Println("修改成功")
			break
		}
		fmt.Println("没有当前学生")
	}
}

// 查询学生信息
func (stuList *StudentList) FindAll() {
	//fmt.Printf("%#v\n %v\n %v\n", stuList, len(stuList.List), cap(stuList.List))
	for _, v := range stuList.List {
		fmt.Printf("ID:%d, NAME:%s, SEX:%s \n", v.Id, v.Name, v.Sex)
	}
}
