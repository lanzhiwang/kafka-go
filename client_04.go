group, err := kafka.NewConsulGroup(ConsulConfig{
	Name: "my-consumer-group",
	Scheme: kafka.Spread,
	Addr: "localhost:8500",
})

readers, err := group.NewReaders(...)

iterators := []kafka.MessageIter{}
for reader := range readers {
	iterators = append(iterators, reader.Read(context.Background(), kafka.Offset(0))
}

iter := kafka.NewMultiIter(iterators)

var msg kafka.Message
for iter.Next(&msg) {
	// ...
}