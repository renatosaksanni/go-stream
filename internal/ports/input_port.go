package ports

// MessageConsumer defines the interface for consuming messages
type MessageConsumer interface {
	Consume(out chan<- interface{}) error
}

// Operator defines the interface for processing data
type Operator interface {
	Apply(in <-chan interface{}) <-chan interface{}
}

// WindowOperator defines the interface for windowing operations
type WindowOperator interface {
	Apply(in <-chan interface{}) <-chan interface{}
}
