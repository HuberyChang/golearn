package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer // 全局的
	logDataChan chan *logData
)

// Init 安利康师傅
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	//fmt.Println("...连接kafka成功！")
	// 初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	return
}

// 给外部暴露一个函数，该函数只把日志发送到一个内部指定的channel中
func SendtoChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// sendToKafka 发送到kafka
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)

		}
	}
}
