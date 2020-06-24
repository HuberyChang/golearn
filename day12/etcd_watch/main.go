package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//etcd watch
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("connect etcd success")
	defer cli.Close()

	// watch
	// 排一个哨兵监视key变化
	ch := cli.Watch(context.Background(), "cao")
	// 从通道尝试获取监视的信息
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
