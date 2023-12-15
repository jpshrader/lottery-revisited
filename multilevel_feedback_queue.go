package main

import (
	"math"
	"math/rand"
)

type multilevelFeedbackQueue[T any] []level[T]

type level[T any] struct {
	Priorty int
	Queue   []T
}

func (l level[T]) draw(constant float64, multiplier float64) (*T, bool) {
	for _, item := range l.Queue {
		if float64(rand.Intn(1000)) <= l.getLikelihood(constant, multiplier) {
			return &item, true
		}
	}
	return nil, false
}

func (l level[T]) getLikelihood(constant float64, multiplier float64) float64 {
	return float64(l.Priorty) + math.Log2(constant*multiplier)
}
