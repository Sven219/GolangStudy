package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "台湾"}
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1)) //s1:[北京 上海 台湾],len:3,cap:3
	s1 = append(s1, "广州")
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1)) //s1:[北京 上海 台湾 广州],len:4,cap:6
	s2 := []string{"西安", "合肥", "郑州", "杭州", "芜湖", "深圳", "苏州", "无锡", "常州"}
	// 切片合并
	s1 = append(s1, s2...)
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1)) //s1:[北京 上海 台湾 广州],len:4,cap:6

}
