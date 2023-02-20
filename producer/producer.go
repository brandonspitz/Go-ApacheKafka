package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), network: "tcp", address: "localhost:9092", topic: "topic_test", partition: 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: []byte("Hello World")})
}