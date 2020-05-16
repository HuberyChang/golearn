package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("==================汉字============")
	// 1、判断字符串中汉字数量
	// 难点是判断一个字符是汉字？
	s1 := "fhajh暗黑结婚登记hah复合接口来说都ds"
	// 1、依次拿到字符串字符
	var count int
	for _, c := range s1 {
		// 2、判断是否是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	// 3、汉字数累加
	fmt.Println(count)
	fmt.Println("==================汉字============")

	fmt.Println("==================英语单词============")
	// how do you do 单词出现次数
	s2 := "how do you do"
	// 把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	// 遍历切片存入到map
	m1 := make(map[string]int, 10)
	for _, c := range s3 {
		fmt.Println(c)
		// 如果原来map中不存在c这个key,那么出现次数=1
		if _, ok := m1[c]; !ok {
			m1[c] = 1
		} else { // 如果原来map中存在c这个key，那么出现次数+1
			m1[c]++
		}
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 累计次数
	fmt.Println("==================英语单词============")

	fmt.Println("==================回文判断============")
	ss := "山西运煤车煤运西山"

	// 把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println("[]rune:", r)

	for i := 0; i < len(r)/2; i++ {
		// 山 ss[0] s[len(ss)-1-0]
		// 西 ss[1] s[len(ss)-1-1]
		// 运 ss[2] s[len(ss)-1-2]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

	fmt.Println("==================回文判断============")

}
