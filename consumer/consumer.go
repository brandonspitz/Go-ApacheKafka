package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), network: "tcp", address: "localhost:9092", topic: "topic_test", partition: 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	batch := conn.ReadBatch(1e3, 1e9) //1e3 means 1000
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
	conn.ReadBatchWith()
}