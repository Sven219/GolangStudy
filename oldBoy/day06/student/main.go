package main

import (
	"fmt"
	"os"
)

// 函数版学生管理系统
// 写一个系统能够查看\新增\删除学生
type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student

// 构造函数
func newStudent(name string, id int64) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 1. 创建一个新学生
	// 2. 追加到 map 中
	var id int64
	var name string
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	newStudent := newStudent(name, id)
	allStudent[id] = newStudent
	fmt.Printf("学生%s已添加成功\n", name)

}

func removeStudent() {
	fmt.Print("请输入要删除的学生的学号：")
	var id int64
	fmt.Scanln(&id)
	delete(allStudent, id)
	fmt.Printf("学生%d已删除成功\n", id)

}
func main() {
	allStudent = make(map[int64]*student, 50) // 初始化，开辟内存空间
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
	1. 查看所有学生
	2. 新增学生
	3. 删除学生
	4. 退出`)
		fmt.Print("请输入选择：")
		// 2. 等待用户选择
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了:%d\n", choice)
		// 3. 执行对应的操作
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			removeStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请检查输入选项是否正确")
		}
	}
}
