package main

import (
	"fmt"
	"os"
)

// 学生管理系统
type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 1. 根据用户的输入内容创建一个新的学生
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)
	newStu := student{
		id:   id,
		name: name,
	}
	// 2. 把新的学生放到 s.allStudent map中
	s.allStudent[id] = newStu

	fmt.Println("学号：%d 姓名：%s 已添加成功", id, name)

}

// 删除学生
func (s studentMgr) delStudent() {
	var id int64
	fmt.Printf("请输入要删除学生的学号：")
	fmt.Scanln(&id)
	value, ok := s.allStudent[id]
	if !ok {
		fmt.Println("该学号不存在学生")
		return
	}
	delete(s.allStudent, id)
	fmt.Printf("学号：%d 姓名：%s 已删除成功\n", value.id, value.name)

}

// 修改学生
func (s studentMgr) updateStudent() {
	// 1. 获取用户输入的学号
	var (
		id   int64
		name string
	)
	fmt.Printf("请输入学号：")
	fmt.Scanln(&id)
	// 2. 展示该学号对应的学生信息，如果没有提示查无此人
	value, ok := s.allStudent[id]
	if !ok {
		fmt.Println("该学号不存在学生，请先添加学生")
		return
	}
	fmt.Printf("您修改的学生原姓名为：%s\n", value.name)
	// 3. 请输入修改后的学生名
	fmt.Printf("请输入修改后的姓名：")
	fmt.Scanln(&name)
	// 4. 更新学生姓名
	// newStu := student{
	// 	id:   id,
	// 	name: name,
	// }
	// s.allStudent[id] = newStu
	value.name = name
	s.allStudent[id] = value
	fmt.Printf("学号：%d 姓名：%s 已修改成功\n", id, name)

}

var smr studentMgr

func showMenu() {

	fmt.Println("-----Welcome sms！-----")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出

	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入
		fmt.Printf("请输入选择的菜单：")
		var opt int
		fmt.Scanln(&opt)
		// fmt.Println("你输入的是：", opt)
		switch opt {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.updateStudent()
		case 4:
			smr.delStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("请输入正确的选项")
		}
	}

}
