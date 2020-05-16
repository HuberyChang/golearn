package mylogger

import (
	"fmt"
	"time"
)

// 在终端写日志相关内容

// ConsoleLogger ...
type ConsoleLogger struct { //标黄是因为作为对外暴露的接口，没添加注释
	Level LogLevel
}

func NewConsoleLogger(levelStr string) ConsoleLogger { //同上
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNum := getNum(3)
		fmt.Printf("[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), logString(lv), funcName, fileName, lineNum, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
