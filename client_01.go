import "github.com/segmentio/kafka-go"

config := kafka.ReaderConfig{
	BrokerAddrs: []string{"localhost:9092"},
	Topic: "events",
	Partition: 0
}

reader, err := kafka.NewReader(config)
if err != nil {
	panic(err)
}

iter := reader.Read(context.Background(), kafka.Offset(0))

var msg kafka.Message
for iter.Next(&msg) {
	fmt.Printf("offset: %d, key: %s, value: %s\n", int64(msg.Offset), string(msg.Key), string(msg.Value))
}


