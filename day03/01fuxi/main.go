package main

import "fmt"

func main() {
	var name string
	name = "理想"
	fmt.Println(name)

	var ages [30]int
	ages = [30]int{1, 3, 5, 7, 9, 11}
	fmt.Println(ages)

	ages2 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(ages2)

	var ages3 = [...]int{1: 100, 20: 200}
	fmt.Println(ages3)

	// 二维数组
	var a1 = [...][2]int{ //只有最外层可以使用...
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a1)

	//数组是值类型
	x := [3]int{1, 2, 3}
	y := x     //把x的值拷贝一份到y
	y[1] = 100 //修改的是y，x不变
	fmt.Println(x)

	fmt.Println(x)
	f1(x) // D盘拷贝一份到F盘，修改的是F盘的副本
	fmt.Println(x)

	// 切片
	var s1 []int
	fmt.Println(s1) //没有分配内存
	fmt.Println(s1 == nil)
	s1 = []int{1, 3, 5} //初始化
	fmt.Println(s1)
	// make初始化,同时分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)

	s3 := make([]int, 0, 4)
	fmt.Println(s3)
	fmt.Println(s3 == nil)


	fmt.Println("================================")
	s4 := []int{1, 2, 3}
	s5 := s4 // s5、s4指针指向同一个地址
	// var ss []int		//没申请内存空间
	var ss = make([]int, 3, 3)
	copy(ss, s5)
	fmt.Println(s5)
	s5[1] = 200
	fmt.Println(s5)
	fmt.Println(s4)
	fmt.Println(ss)
	fmt.Println("================================")

	// var s6 []int //nil
	// s6[0] = 100  //尚未分配内存，报错
	// fmt.Println(s6)

	// 切片的扩容策略
	// 如果申请容量大于原来2倍，那就直接扩容至新申请的容量；如果小于1024，就2倍；如果大于1024，就1.25倍；
	var s7 []int
	s7 = append(s7, 1) //自动初始化切片
	fmt.Println(s7)





	fmt.Println("=============指针=================")
	// go语言指针只能读，无法修改，不能修改指针变量指向的地址
	addr := "shahe"
	addrp := &addr
	fmt.Println(addrp)
	fmt.Printf("%T\n", addrp)
	addrv := *addrp //根据内存地址找值
	fmt.Println(addrv)
	fmt.Println("=============指针=================")


	fmt.Println("=============Map=================")
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["hahh"] = 200
	// m1["hahj"] = 100
	fmt.Println(m1)
	fmt.Println(m1["hahj"]) 	// 如果key不在，value返回定义时对应类型的0值

	// 如果返回值是布尔值，我们通常用ok去接收
	score, ok := m1["hahj"]
	if !ok {
		fmt.Println("不存在")
	}else{
		fmt.Println("分数是：", score)
	}
	//删除
	delete(m1,"lixiang")	// 删除的key不存在，什么也不干
	delete(m1,"hahh")
	fmt.Println(m1)
	fmt.Println(m1 == nil)	//已经申请了内存空间
	fmt.Println("=============Map=================")

}

func f1(a [3]int) {
	//go语言的函数传参传递的是值（Ctrl+C Ctrl+V）
	a[1] = 100 //此处修改的是副本的值
}
