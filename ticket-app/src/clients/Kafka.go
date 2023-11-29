package clients

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers: []string {"kafka:29092"},
		Topic: "payment",
		GroupID: "g1",
		MaxBytes: 10,
	} 

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occured while consume kafka event", err)
			continue
		}
		fmt.Println("Message is: ", string(m.Value));
	}
}