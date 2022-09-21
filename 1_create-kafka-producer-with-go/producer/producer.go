package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
conn, _ :=	kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "ContohTopic", 0)

conn.SetWriteDeadline(time.Now().Add(time.Second*10)) // if something went wrong, execute deadline and write error msg

conn.WriteMessages(kafka.Message{Value: []byte("hello kafka, im from golang!")})
}