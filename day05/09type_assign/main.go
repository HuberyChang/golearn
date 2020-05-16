package main

import "fmt"

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("type:%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("wrong!")
	} else {
		fmt.Println("ture:", str)
	}
	fmt.Println(str)
}

// 类型断言2
func assign1(a interface{}) {
	fmt.Printf("type:%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	case int64:
		fmt.Println("int64", t)
	}
}

func main() {
	assign(100)
	assign1("ifpjpi")
}
