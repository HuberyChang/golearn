package taillog

import "golearn/logs/etcd"

var tskMgr *taillogMgr

type taillogMgr struct {
	logEntry []*etcd.LogEntry
	tskMap   map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry: logEntryConf, // 把当前的日志收集项配置信息保存起来
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
