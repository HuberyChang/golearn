package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
	map属于引用类型
	存储特点：
		A：无序
		B：key不可重复，否则新key会覆盖原来的key，程序不报错

	语法结构：
		1、创建map
			var map1 map[key类型]value类型
				nil map，无法直接使用

			var map2 =make(map[key类型]value类型)

			var map3 =map[key类型]value类型{key:value,key:value,key:value...}
		2、添加/修改
			map[key]=value
		3、获取
			map[key]=value
		4、删除数据
			delete()

	每种数据类型默认值：
		int:0
		float:0.0
		string:""
		array:[00000]

		slice:nil
		map:nil
*/

func main() {
	// 1.创建
	var map1 map[int]string         //没有初始化，nil
	var map2 = make(map[int]string) //创建
	var map3 = map[string]int{"GO": 90, "java": 80, "php": 98}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	//2.nil map
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	//3.存储数据
	map1[1] = "hel"
	map1[2] = "qour"
	map1[3] = "109u"
	map1[4] = "dkfq"
	map1[5] = "多了几分去啊"
	map1[6] = "fqj"
	map1[7] = ""

	// 4.获取数据
	//根据key获取对应的value，如果不存在这个key，拿到对应类型的“0”值
	fmt.Println(map1)
	fmt.Println(map1[2])
	fmt.Println(map1[20]) // ""
	// 约定俗成，用“ok”接收返回的bool值
	value, ok := map1[7]
	if ok {
		fmt.Printf("key存在：%v\n", value)
	} else {
		fmt.Println("key不存在")
	}

	// 5.修改数据
	map1[7] = "djf[q9i打分卡"
	fmt.Println(map1[7])

	// 6. 删除
	delete(map1, 7)
	fmt.Println(map1)

	// 7.长度
	fmt.Println(len(map1))
	fmt.Println("===============================")

	/*
		1.获取所有的key--->切片/数组
		2.进行排序
		3.遍历key--->map[key]
	*/
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string]int, 200) // 初始化随机种子
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(10)           // 生成0-99的随机整数
		scoreMap[key] = value
	}

	// 取出map中的所有key存入切片
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
