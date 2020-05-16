package main

import "fmt"

// 需指定数组的类型和长度
func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true false true false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)		//a1:[3]bool a2:[4]bool

	// 数组初始化
	// 不初始化：默认元素都是零，(布尔值：false；整型、浮点型都是0,；字符串："")
	fmt.Println(a1, a2)		//[false false false] [false false false false]

	// 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)		//[true true true]
	// 方式2 根据初始值自动推断长度
	a100 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a100)	//[0 1 2 3 4 5 6 7 8 9]
	// 方式3
	a3 := [5]int{1, 2}
	fmt.Println(a3)		//[1 2 0 0 0]
	// 根据索引初始化值
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4) 	//[1 0 0 0 2]

	// 数组遍历
	citys := [...]string{"北京", "上海", "成都"}
	// 根据索引
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])	//北京
								//上海
								//成都
	}
	// 根据range
	for i, v := range citys {
		fmt.Println(i, v)	//0 北京
							//1 上海
							//2 成都
	}

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{12, 13},
		[2]int{14, 15},
		[2]int{16, 17},
	}
	fmt.Println(a11)	//[[12 13] [14 15] [16 17]]

	// 多维数组的遍历
	for i, v1 := range a11 {
		fmt.Println(i, v1)
		for i, v2 := range v1 {
			fmt.Println(i, v2)	// 0 [12 13]
								// 0 12
								// 1 13
								// 1 [14 15]
								// 0 14
								// 1 15
								// 2 [16 17]
								// 0 16
								// 1 17
		}
	}





	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)	// [1 2 3] [100 2 3]


	//
	c := [...]int{1, 2, 3, 4, 5}
	sum := 0
	for  _, v := range c{	// 丢掉索引
		sum += v
	}
	fmt.Println(sum) 	// 15


	// 找出和为6的两个元素的下标
	// 两个for循环，第一个for循环从外层开始遍历
	// 第二个for循环
	for i := 0; i < len(c); i++{
		for j := i + 1; j < len(c); j++{
			if c[i] + c[j] == 6{
				fmt.Println(i, j)	// 0 4
									// 1 3
			}

		}
	}
}
