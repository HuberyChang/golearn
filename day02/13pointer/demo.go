package main

import "fmt"

// 房子里的人变了
func swap(a, b *int) { // 类似于小明，小王住的房子分别编号为A，B；现在把小明、小王从各自的房子里互换，现在A、B两套房子住的人就是互换了人之后的
	//t := *a
	//*a = *b
	//*b = t
	*a, *b = *b, *a
}

// 房子里边的人没变
func swap1(a, b *int) { // 类似于小明，小王住的房子分别编号为A，B；现在把房子的编号互换，本质小明，小王还是住的他们原来的房子，只不过房子的名字互换一下
	//fmt.Println(a, b)
	//fmt.Println(*a, *b)
	b, a = a, b
	//fmt.Println(a, b)
	//fmt.Println(*a, *b)
}

func main() {
	x, y := 1, 2
	//fmt.Printf("%p,%p=====1\n", &x, &y)
	//swap(&x, &y)
	pa, pb := &x, &y
	fmt.Printf("%p,%p====1\n", pa, pb)
	fmt.Printf("%p,%p====2\n", &pa, &pb)
	//swap(pa, pb)
	swap1(pa, pb)
	//fmt.Printf("%p,%p=====2\n", &x, &y)
	fmt.Printf("%p,%p====3\n", pa, pb)
	fmt.Printf("%p,%p====4\n", &pa, &pb)
	fmt.Println(x, y)

	/*	var a = 10
		var p *int = &a
		a = 100
		fmt.Println(a)
		//fmt.Println(&a)
		fmt.Println(p)
		fmt.Println(&p)
		*p = 250
		fmt.Println(p)
		fmt.Println(&p)
		fmt.Println(a)
		fmt.Println(*p)*/

	/*	for i, i2 := range []int{1, 2, 3, 4} {
			fmt.Printf("key:%d value:%d\n", i, i2)
		}

		var str = "hello 哈哈"
		for i, i2 := range str {
			fmt.Printf("key:%v value:%q\n", i, i2)

		}

		m := map[string]int{
			"hello": 100,
			"world": 200,
		}
		for i, i2 := range m {
			fmt.Println(i, i2)
		}

		c := make(chan int)
		go func() {
			c <- 1
			c <- 2
			c <- 3
			close(c)
		}()
		for i2 := range c {
			fmt.Println(i2)
		}*/
}
