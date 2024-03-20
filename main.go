package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mqClient := &MqClient{
		"localhost:8081",
		"group_id",
		"topic_id",
		nil,
	}

	err := mqClient.InitAndSubscribe()
	if err != nil {
		panic(err)
	}
	defer mqClient.Close()

	// 创建信号处理器，捕获中断信号以关闭程序
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 循环消费消息
ConsumerLoop:
	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			break ConsumerLoop
		default:
			msg := mqClient.GetTopic()
			if msg == "" {
				fmt.Println("continue")
				continue
			} else {
				task := Task{}
				if err := task.Parse(msg); err != nil {
					fmt.Println("Parse err")
					continue
				}
				if err := task.Do(); err != nil {
					fmt.Println("Do err")
				}

			}

		}
	}

}
