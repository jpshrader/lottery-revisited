package main

import (
	"math"
	"math/rand"
)

type multilevelFeedbackQueue []level

type level struct {
	Queue []participant
}

func (l level) draw(constant float64, multiplier float64) (*participant, bool) {
	for _, item := range l.Queue {
		if float64(rand.Intn(1000)) <= l.getLikelihood(item.Priority, constant, multiplier) {
			return &item, true
		}
	}
	return nil, false
}

func (l level) getLikelihood(priority int, constant float64, multiplier float64) float64 {
	return constant + math.Log2(float64(priority+1)*multiplier)
}
