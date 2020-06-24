package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方读
		MustExist: false,                                // 必须存在，不在就报错
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed ,err:", err)
		return
	}
	return
}

// ReadChan 读日志
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
