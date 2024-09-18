package processing

import (
	"go-stream/internal/ports"
)

type MapFunction func(interface{}) (interface{}, error)

type MapOperator struct {
	Fn MapFunction
}

func NewMapOperator(fn MapFunction) ports.Operator {
	return &MapOperator{Fn: fn}
}

func (op *MapOperator) Apply(in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for data := range in {
			result, err := op.Fn(data)
			if err != nil {
				// Handle error (e.g., log and continue)
				continue
			}
			out <- result
		}
	}()
	return out
}
