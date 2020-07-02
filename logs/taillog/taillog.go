package taillog

import (
	"fmt"
	"golearn/logs/kafka"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

// TailTask : 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件哪个地方读
		MustExist: false,                                // 必须存在，不在就报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed ,err:", err)

	}
	go t.run() // 直接采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			// 先把日志信息发送到一个通道中
			kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
		}
	}
}
