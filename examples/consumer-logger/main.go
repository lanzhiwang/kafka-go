package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
	scram "github.com/segmentio/kafka-go/sasl/scram"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	mechanism, err := scram.Mechanism(scram.SHA512, "admin", "mTtCGCWKZoA_xRzW")
	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   dialer,
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL := "10.0.128.237:30777"
	topic := "my-topic"
	groupID := "my-group-1"

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
