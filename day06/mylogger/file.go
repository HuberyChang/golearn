package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger ...
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *LogMsg
}

// LogMsg ...
type LogMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger ... 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *LogMsg, 50000),
	}
	err = fl.initFile() //按照文件名和文件路径将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启一个后台goroutine去往文件写日志
	for i := 1; i < 5; i++ {
		go f.writeLogBackground()
	}
	return nil
}

// 判断是否需要记录该日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

// 判断文件是否需要切割
func (f *FileLogger) checkLogSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file size failed,err:%v", err)
		return false
	}
	// 如果当前文件大小大于等于文件大小的最大值，就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file failede, err:%v", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件的完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //	拼接一个日志文件备份的名字
	// 1. 关闭当前文件
	f.fileObj.Close()
	// 2. 备份一下日志，rename
	os.Rename(logName, newLogName)
	// 3. 打开一个新文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {

	for {
		if f.checkLogSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s-%s-%d] %s\n", logTmp.timestamp, logString(logTmp.level), logTmp.funcName, logTmp.fileName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkLogSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.fileObj = newFile
				}
			}
			// 如果要记录的日志级别大于等于error，还要在error日志文件中保存一份
			fmt.Fprintf(f.errFileObj, logInfo)
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNum := getNum(3)
		// 先把日志发送到通道中
		// 1. 造一个logMsg对象
		logTmp := &LogMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			line:      lineNum,
			timestamp: now.Format("2006-01-02 15:04:05"),
		}
		select {
		case f.logChan <- logTmp:
		default:
		}
	}

}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	// fmt.Println(msg)

	f.log(DEBUG, format, a...)

	// if l.enable(DEBUG) {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }

	// if l.Level > DEBUG {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	// fmt.Println(msg)
	f.log(INFO, format, a...)
	// if l.enable(INFO) {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	// fmt.Println(msg)
	f.log(WARNING, format, a...)
	// if l.enable(WARNING) {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	// fmt.Println(msg)
	f.log(ERROR, format, a...)
	// if l.enable(ERROR) {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	// fmt.Println(msg)
	f.log(FATAL, format, a...)
	// if l.enable(FATAL) {
	// 	now := time.Now()
	// 	fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	// }
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
