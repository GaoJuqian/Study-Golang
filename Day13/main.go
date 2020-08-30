package main

import (
	"encoding/json"
	"fmt"
)

// 学生列表
type studentList struct {
	Name    string    `json:"name"` // 大写 注意对外部包的可见性
	Student []student `json:"student"`
}

// 学生信息
type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 学生构造函数
func newStudent(id int, name string) *student {
	return &student{
		ID:   id,
		Name: name,
	}
}

func main() {

	// 初始化
	test1 := &studentList{
		Name:    "学生XX列表",
		Student: make([]student, 0, 50),
	}

	// 添加数据
	for i := 1; i < 50; i++ {
		stu := newStudent(i, fmt.Sprintf("stu%02d", i))
		test1.Student = append(test1.Student, *stu)
	}
	//	json格式化
	studentListJSON, err := json.Marshal(test1)
	if err != nil {
		fmt.Printf("json marshal err:%v\n", err)
	}
	// 返回[]byte
	fmt.Printf("%s\n", studentListJSON)

	// 转换为字符串
	test1Str := fmt.Sprintf("%s", studentListJSON)

	// json字符串反转为类型
	var test2 studentList
	err = json.Unmarshal([]byte(test1Str), &test2) // 反序列化 传接受的类型地址
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err)
	}
	fmt.Printf("%v\n", test2)
}
