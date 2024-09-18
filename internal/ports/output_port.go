package ports

// MessageProducer defines the interface for producing messages
type MessageProducer interface {
	Produce(in <-chan interface{}) error
}
