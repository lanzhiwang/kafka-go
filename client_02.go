import "github.com/segmentio/kafka-go"

group, err := kafka.NewConsulGroup(ConsulConfig{
	Name: "my-consumer-group",
	Addr: "localhost:8500",
})

reader, err := group.NewReader(ReaderConfig{
	BrokerAddrs: []string{"localhost:9092"},
	Topic: "events",
	Partitions: 100,
})
