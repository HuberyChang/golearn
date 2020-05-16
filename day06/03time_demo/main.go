package main

import (
	"fmt"
	"time"
)

// 时间

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// time.Unix
	ret := time.Unix(1588511768, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	// 时间间隔
	fmt.Println(time.Second)

	fmt.Println(now.Add(24 * time.Hour))

	// 时间间隔
	nextYear, err := time.Parse("2006-01-02", "2020-05-02")
	if err != nil {
		fmt.Println("failed！", err)
		return
	}
	d := now.UTC().Sub(nextYear)
	fmt.Println(d)
	fmt.Println("+++++++++++++")

	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 格式化时间
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))

	// sleep
	n := 5
	fmt.Println("开始sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("s秒钟过去了")

}
