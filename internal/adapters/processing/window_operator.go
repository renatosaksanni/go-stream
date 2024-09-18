package processing

import (
	"go-stream/internal/ports"
	"time"
)

type TumblingWindowOperator struct {
	Duration      time.Duration
	AggregateFunc func([]interface{}) (interface{}, error)
}

func NewTumblingWindowOperator(duration time.Duration, aggFunc func([]interface{}) (interface{}, error)) ports.WindowOperator {
	return &TumblingWindowOperator{
		Duration:      duration,
		AggregateFunc: aggFunc,
	}
}

func (w *TumblingWindowOperator) Apply(in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		var windowData []interface{}
		timer := time.NewTimer(w.Duration)

		for {
			select {
			case data, ok := <-in:
				if !ok {
					// Input channel closed
					if len(windowData) > 0 {
						aggregatedData, err := w.AggregateFunc(windowData)
						if err == nil {
							out <- aggregatedData
						}
					}
					return
				}
				windowData = append(windowData, data)
			case <-timer.C:
				if len(windowData) > 0 {
					aggregatedData, err := w.AggregateFunc(windowData)
					if err == nil {
						out <- aggregatedData
					}
					windowData = nil
				}
				timer.Reset(w.Duration)
			}
		}
	}()
	return out
}
