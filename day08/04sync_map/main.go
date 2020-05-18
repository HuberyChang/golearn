package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go内置的map不是并发安全

// var m = make(map[string]int)
var m = sync.Map{}
var lock sync.Mutex

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int) {
// 	m[key] = value
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 30; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			lock.Lock()
// 			set(key, n)
// 			lock.Unlock()
// 			fmt.Printf("k=:%v v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// lock.Lock()
			m.Store(key, n) // 必须使用sync.map的Store方法存键值对
			// lock.Unlock()
			value, _ := m.Load(key) // 必须使用sync的load方法根据key取值
			fmt.Printf("k=:%v v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
