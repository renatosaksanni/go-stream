package application

import (
	"go-stream/internal/ports"
)

type StreamService struct {
	InputPort  ports.MessageConsumer
	OutputPort ports.MessageProducer
	Operators  []ports.Operator
	Window     ports.WindowOperator
}

func NewStreamService(input ports.MessageConsumer, output ports.MessageProducer, ops []ports.Operator, window ports.WindowOperator) *StreamService {
	return &StreamService{
		InputPort:  input,
		OutputPort: output,
		Operators:  ops,
		Window:     window,
	}
}

func (s *StreamService) Start() error {
	dataCh := make(chan interface{})
	errCh := make(chan error)

	// Start consuming messages
	go func() {
		err := s.InputPort.Consume(dataCh)
		if err != nil {
			errCh <- err
		}
	}()

	// Process messages
	processedCh := s.process(dataCh)

	// Handle errors
	select {
	case err := <-errCh:
		return err
	default:
		// No error, continue
	}

	// Start producing messages
	err := s.OutputPort.Produce(processedCh)
	if err != nil {
		return err
	}

	return nil
}

func (s *StreamService) process(dataCh <-chan interface{}) <-chan interface{} {
	var ch <-chan interface{} = dataCh

	// Apply operators
	for _, op := range s.Operators {
		ch = op.Apply(ch)
	}

	// Apply windowing if set
	if s.Window != nil {
		ch = s.Window.Apply(ch)
	}

	return ch
}
