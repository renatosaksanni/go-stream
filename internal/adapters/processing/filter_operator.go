package processing

import (
	"go-stream/internal/ports"
)

type FilterFunction func(interface{}) bool

type FilterOperator struct {
	Fn FilterFunction
}

func NewFilterOperator(fn FilterFunction) ports.Operator {
	return &FilterOperator{Fn: fn}
}

func (op *FilterOperator) Apply(in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for data := range in {
			if op.Fn(data) {
				out <- data
			}
		}
	}()
	return out
}
