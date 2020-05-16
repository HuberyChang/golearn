package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // true，表示还没在内存开辟空间
	m1 = make(map[string]int, 10) // 提前估算还容量，避免在运行期间动态扩展
	m1["lei"] = 18
	m1["ha"] = 35
	fmt.Println(m1)
	fmt.Println(m1["jian"]) // 如果不存在这个key，拿到对应类型的“0”值
	fmt.Println(m1["ha"])

	// 约定俗成，用“ok”接收返回的bool值
	// value, ok := m1["hah"]
	// if !ok {
	// 	fmt.Println("no the key")
	// }else{
	// 	fmt.Println(value)
	// }

	// map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 遍历k
	for k := range m1 {
		fmt.Println(k)
	}
	// 遍历v
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "lei")
	fmt.Println(m1)
	delete(m1, "hah")
	fmt.Println(m1)

	//
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string]int, 200)	// 初始化随机种子
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)	// 生成stu开头的字符串
		value := rand.Intn(100) 	// 生成0-99的随机整数
		scoreMap[key] = value
	}

	// 取出map中的所有key存入切片
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
