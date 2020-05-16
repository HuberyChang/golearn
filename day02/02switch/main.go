package main

import "fmt"

func main() {
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	}else if n == 2 {
		fmt.Println("食指")
	}else if n == 3 {
		fmt.Println("中指")
	}else if n == 4 {
		fmt.Println("无名指")
	}else if n == 5 {
		fmt.Println("小指")
	}else {
		fmt.Println("无效字符")
	}


	// switch
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default :
		fmt.Println("错")
	}
}
