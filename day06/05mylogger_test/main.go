package main

import (
	"golearn/day06/mylogger"
	"time"
)

var log mylogger.Logger //声明一个全局的变量

// 测试自己写的日志库
func main() {
	log = mylogger.NewConsoleLogger("debug")                         //console级别的日志，输出到终端
	log = mylogger.NewFileLogger("debug", "./", "test.log", 10*1024) //文件类型的日志，输出到文件

	// log = mylogger.NewLogger()
	for {
		log.Debug("debug test...")
		log.Info("info test...")
		log.Warning("warning test...")
		id := 10010
		name := "理想"
		log.Error("error test...,id:%s,name:%s", id, name)
		log.Fatal("fatal test...")
		time.Sleep(time.Second * 3)
	}
}
