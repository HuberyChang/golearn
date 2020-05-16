package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数据
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Printf("err:%s", err)
		return
	}
	fmt.Println(ret1)

	// 简单字符串解析出对应的数据
	ret, _ := strconv.Atoi(str)
	fmt.Println(ret)

	// 把数字转换为字符串类型
	i := 97
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)

	// 简单数字转字符串
	sret := strconv.Itoa(i)
	fmt.Printf("%#v\n", sret)

	// 从字符串中解析出bool值
	boolstr := "true"
	boolValue, _ := strconv.ParseBool(boolstr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串中解析float
	floatStr := "1.1289347809"
	floatValue, _ := strconv.ParseFloat(floatStr, 18)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

}
