package main

type lottery[T any] struct {
	MultilevelFeedbackQueue multilevelFeedbackQueue[T]
	Constant                float64
	Multiplier              float64
}

func (l lottery[T]) draw() (*T, bool) {
	for true {
		if item, won := l.drawWorker(); won {
			return item, true
		}
	}
	return nil, false
}

func (l lottery[T]) drawWorker() (*T, bool) {
	for _, level := range l.MultilevelFeedbackQueue {
		if item, won := level.draw(l.Constant, l.Multiplier); won {
			return item, true
		}
	}
	return nil, false
}
