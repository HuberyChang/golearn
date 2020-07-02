package main

import "fmt"

// map和slice组合
func main() {
	// 元素类型为map的切片
	// 对map和slice同时初始化
	map1 := map[string]string{"name": "大开发", "age": "30", "sex": "na", "addr": "挨决独守空房"}
	map2 := make(map[string]string)
	map2["name"] = "挨决看到了"
	map2["age"] = "23"
	map2["sex"] = "nv"
	map2["addr"] = "对方情况我确认"
	map3 := map[string]string{"name": "的罚款", "age": "989", "sex": "asd", "addr": "拉手孔江东父老"}
	//var s1 = make([]map[int]string, 1, 10) // 初始化slice
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	for i, val := range s1 {
		fmt.Printf("第%d个人的信息：%v\n", i+1, val)
		//fmt.Printf("%t")
	}
	// 没对内部的map做初始化
	// s1[0][100] = "A"
	//s1[0] = make(map[int]string, 1) // 初始化map
	//s1[0][100] = "hang"
	//fmt.Println(s1)

	// 值为slice类型的map
	var m1 = make(map[string][]int, 10)
	m1["沙河"] = []int{10, 20, 30}
	fmt.Println(m1)
}
