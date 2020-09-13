package goBreaker

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker()
	cb.State=2
	cb.HalfToClose = DefaultToClose
	cb.CloseToOpen = DefaultToOpen
	cb.HalfToOpen = DefaultToOpen

	cb.Execute(func() (interface{}, error) {
		i := rand.Intn(2)
		if i == 0 {
			return nil, errors.New("error")
		}
		return nil, nil
	})
	fmt.Println(cb.State)
}
