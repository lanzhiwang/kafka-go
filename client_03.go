import "github.com/segmentio/kafka-go"

group, err := kafka.NewConsulGroup(ConsulConfig{
	Name: "my-consumer-group",
	Addr: "localhost:8500",
})

readers := []kafka.Reader{}
iterators := []kakfa.MessageIter{}

for i := 0; i < 100; i++ {
	reader, err := group.NewReader(ReaderConfig{
		BrokerAddrs: []string{"localhost:9092"},
		Topic: "events",
		Partitions: 100,
	})

	iter := reader.Read(context.Background(), kafka.Offset(0))

	readers = append(readers, reader)
	iterators = append(iterators, iter)
}

iter := kafka.NewMultiIter(iterators)

var msg kafka.Message
for iter.Next(&msg) {
	// ...
}