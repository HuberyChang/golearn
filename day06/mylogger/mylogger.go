package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// LogLevel ...
type LogLevel uint16

// Logger ...
type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	// UKNOWN ...
	UKNOWN LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(str string) (LogLevel, error) {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		// error:=fmt.Errorf("")
		error := errors.New("无效日志级别")
		return UKNOWN, error
	}
}

func logString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "IINFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		break
	}
	return "debug"
}

func getNum(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return

}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...) // 加...是把a这个切片的元素当成一个整体整个传进去
	// msg := fmt.Sprintf(format, a) 	// 不加...，是把a这个切片的所有元素拿出来一个一个传进去
	now := time.Now()
	funcName, fileName, lineNum := getNum(3)
	fmt.Printf("[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), logString(lv), funcName, fileName, lineNum, msg)
}
