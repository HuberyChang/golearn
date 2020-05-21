package main

import (
	"fmt"
	"golearn/day09/split_string"
)

func main() {
	var ret []string
	ret = split_string.Split("abcdbef", "b")
	fmt.Printf("%v\n", ret)
	ret = split_string.Split("abcdbef", "a")
	fmt.Printf("%v\n", ret)
	ret = split_string.Split("abcdbef", "c")
	fmt.Printf("%v\n", ret)
	ret = split_string.Split("abcdbef", "d")
	fmt.Printf("%v\n", ret)
	ret = split_string.Split("abcdbef", "e")
	fmt.Printf("%v\n", ret)
}
