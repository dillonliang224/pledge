package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader: ", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)

	b := make([]byte, 10e3)
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch: ", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection: ", err)
	}
}
