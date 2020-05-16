package main

import "fmt"

func main() {
	s := "hello沙河"
	// len()求得是byte字节数量
	n := len(s) // 求字符串s的长度，把长度保存在变量n中
	fmt.Println(n)

	/*	Go 语言的字符有以下两种：

		 	uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
			rune类型，代表一个 UTF-8字符。
			当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32
	*/

	// 遍历字符串
	// for i := 0; i < len(s); i++{
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n",s[i])  // %c:字符
	// }

	// for _, c := range s{ // 从字符串中拿出具体的字符
	// 	fmt.Printf("%c\n",c)
	// 	fmt.Println(c)
	// }

	// 字符串修改
	s2 := "白萝卜"      // [白 萝 卜]
	s3 := []rune(s2) // 把字符串强制转化为一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换为字符串

	c1 := "红"
	c2 := '红' // rune(int32) 一个汉字三个字节，utf8带一个前缀，共4个字节，4*8=32
	fmt.Printf("c1:%T,c2:%T", c1, c2)

	n1 := 10
	f := float64(n1)
	fmt.Printf("%T\n", f)
}
