package main

import "fmt"

func main() {
	// s := "hello 沙河"
	// for i, v := range s {
	// 	fmt.Printf("%d %c\n", i, v)
	// }
	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}
}
