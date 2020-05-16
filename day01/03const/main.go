package main

import "fmt"

// 常量
// 常量定义完不可再修改，程序运行期间不会改变的量
const pi = 3.1415926

const(
	statusOk = 200
	notFound = 404
)

// 批量声明变量，某一行没写等号，默认与上一行一致
const(
	n1 = 100
	n2       // 默认n2 = 100
	n3
)

// iota
// iota在const关键字出现时被置为0。const中，每增加一行常量声明，将使 iota 计数一次
const(
	n4 = iota   // n4=0
	n40         // 1
	n41			// 2
	n42			// 3
)


const(
	b1 = iota	// 0
	b2 = iota	// 1
	_  			// 2  等价于 _ = iota，使用'_'跳过某些值
	b3			// 3
)


// 插队
const(
	c1 = iota	// 0
	c2 = 100	// 100
	c3			// 100
	c4 = iota	// 3  
	c5			// 4
)

const(
	d1,d2 = iota + 1, iota + 2	//d1:1 d2:2
	d3,d4 = iota + 1, iota + 2	//d3:2 d4:3
)

// 定义数量级
const(
	_ = iota	//跳过0
	KB = 1 << (10 * iota)	// iota = 1，1左移10位
	MB = 1 << (10 * iota)	// iota = 2，1左移20位
	GB = 1 << (10 * iota)	// iota = 3，1左移30位
	TB = 1 << (10 * iota)	// iota = 4，1左移40位
	PB = 1 << (10 * iota)	// iota = 5，1左移50位
)

func main()  {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(n4)
	fmt.Println(n41)
}

