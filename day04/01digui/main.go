package main

import "fmt"

//递归：自己调用自己
//处理相同问题，但是问题规模越来越小
//递归一定要有明确退出条件

//上楼梯：n个台阶，一次可以走1步，也可以走2步，有多少种走法。
func ft(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ft(n-1) + ft(n-2)
}

// func f(n uint64) uint64 {
// 	if n <= 1 {
// 		return 1
// 	}
// 	return n * f(n-1)
// }

func main() {
	// ret := f(6)
	// fmt.Println(ret)
	ret := ft(4)
	fmt.Println(ret)
}
