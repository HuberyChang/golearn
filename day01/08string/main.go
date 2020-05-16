package main

import (
	"fmt"
	"strings"
)

func main()  {
	path := "\"E:\\goStudy\\golearn\\day01\""
	fmt.Printf(path)

	s1 := `fh;aj
	fa;lfa
	jf;al;kds;576jf`
	fmt.Println(s1)
	// 字符串相关操作
	fmt.Println(len(s1))

	// 字符串拼接
	name := "jhdas"
	world := "ha"
	s2 := name + world
	// fmt.Println(name + world)
	fmt.Println(s2)
	s3 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(s3)
	fmt.Printf("%s%s", name, world)

	// 分割
	ret := strings.Split(s1, ";")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s2,"ha"))

	// 前缀
	fmt.Println(strings.HasPrefix(s2,"jh"))
	// 后缀
	fmt.Println(strings.HasSuffix(s2,"ha"))

	// 索引
	s4 := "abcdefb"
	fmt.Println(strings.Index(s4,"b"))
	fmt.Println(strings.LastIndex(s4,"b"))

	// 拼接
	fmt.Println(strings.Join(ret,"+"))

}