package main

import (
	"fmt"
	"golearn/logs/conf"
	"golearn/logs/etcd"
	"golearn/logs/kafka"
	"golearn/logs/taillog"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

/* 注释1 */
//func run() {
//	// 1. 读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			// 2. 发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

func main() {
	// 0.加载配置文件
	//cfg, err := ini.Load("./conf/config.ini")
	//fmt.Println(cfg.Section("kafka").Key("address"))
	//fmt.Println(cfg.Section("kafka").Key("topic"))
	//fmt.Println(cfg.Section("taillog").Key("path"))

	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}

	/* 注释1 */
	//fmt.Println(cfg.KafkaConf.Address)
	//fmt.Println(cfg.KafkaConf.Topic)
	//fmt.Println(cfg.TaillogConf.FileName)

	//1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	//2.初始化etcd连接
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")

	// 2.1 从etcd中拉取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	// 2.2 排一个哨兵监视收集项的变化
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success,%v\n", logEntryConf)
	for inx, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", inx, value)
	}
	//3.收集日志发往kafka
	//3.1 循环每一个日志，创建tailObj
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		err = taillog.Init(logEntry.Path)
		if err != nil {
			fmt.Printf("Init taillog failed,err:%v\n", err)
			return
		}
	}
	//3.2 发往kafka

	//2.打开日志文件准备收集日志
	/* 注释1 */
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Printf("init taillog failed:%v\n", err)
	//	return
	//
	//}
	//fmt.Println("init taillog success")
	//run()
}
