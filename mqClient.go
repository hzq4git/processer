package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type MqClient struct {
	host          string
	gp_id         string
	topic         string
	consumer_impl *kafka.Reader
}

func (mqClient *MqClient) InitAndSubscribe() error {
	config := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  mqClient.gp_id,
		Topic:    mqClient.topic,
		MaxBytes: 10e6, // 10MB
	}
	consumer := kafka.NewReader(config)
	mqClient.consumer_impl = consumer
	return nil
}

func (mqClient *MqClient) GetTopic() (msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	ev, err := mqClient.consumer_impl.ReadMessage(ctx)
	if err == nil {
		msg = string(ev.Value)
	}

	return
}

func (mqClient *MqClient) Close() {
	mqClient.consumer_impl.Close()
}
