package main

type lottery struct {
	MultilevelFeedbackQueue multilevelFeedbackQueue
	Constant                float64
	Multiplier              float64
}

func (l lottery) draw() (*participant, bool) {
	for {
		if item, won := l.drawWorker(); won {
			return item, true
		}
	}
}

func (l lottery) drawWorker() (*participant, bool) {
	for _, level := range l.MultilevelFeedbackQueue {
		if item, won := level.draw(l.Constant, l.Multiplier); won {
			return item, true
		}
	}
	return nil, false
}
