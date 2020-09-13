package goBreaker

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

func (c *Counts) OnRequest() {
	atomic.AddUint32(&c.Requests, 1)
}

func (c *Counts) OnSuccess() {
	atomic.AddUint32(&c.TotalSuccesses, 1)
	atomic.AddUint32(&c.ConsecutiveSuccesses, 1)
	atomic.StoreUint32(&c.ConsecutiveFailures, 0)
}

func (c *Counts) OnFailure() {
	atomic.AddUint32(&c.TotalFailures, 1)
	atomic.AddUint32(&c.ConsecutiveFailures, 1)
	atomic.StoreUint32(&c.ConsecutiveSuccesses, 0)
}

func (c *Counts) Reset() {
	atomic.StoreUint32(&c.Requests, 0)
	atomic.StoreUint32(&c.TotalSuccesses, 0)
	atomic.StoreUint32(&c.ConsecutiveSuccesses, 0)
	atomic.StoreUint32(&c.TotalFailures, 0)
	atomic.StoreUint32(&c.ConsecutiveFailures, 0)
}

type State int

const (
	CloseState = iota
	HalfOpenState
	OpenState
)

var OpenStateError = errors.New("断路器正在开启中")

type ReadyToTrip func(counts Counts) bool

type CircuitBreaker struct {
	State  State
	Counts Counts

	Mutex sync.Mutex

	CloseToOpen ReadyToTrip
	HalfToOpen  ReadyToTrip
	HalfToClose ReadyToTrip

	TimeOut    time.Duration
	expireTime time.Time
}

func NewCircuitBreaker() *CircuitBreaker{
	cb := new(CircuitBreaker)
	cb.Mutex = sync.Mutex{}
	cb.TimeOut = 10*time.Millisecond
	cb.expireTime = time.Now()
	return cb
}

func (cb *CircuitBreaker) beforeRequest() error {
	if cb.State == OpenState {
		if cb.expireTime.Before(time.Now()) {
			cb.nextState(HalfOpenState)
			return nil
		}
		return OpenStateError
	}
	return nil
}

func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	if err := cb.beforeRequest(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			cb.afterRequest(false)
			return
		}
	}()
	cb.Counts.OnRequest()
	resp, err := req()
	cb.afterRequest(err == nil)
	return resp, nil
}

func (cb *CircuitBreaker) afterRequest(success bool) {
	if !success {
		cb.Counts.OnFailure()
		if cb.CloseToOpen(cb.Counts) {
			cb.nextState(OpenState)
		}
	} else {
		cb.Counts.OnSuccess()
		if cb.State == HalfOpenState {
			if cb.HalfToClose(cb.Counts) {
				cb.nextState(CloseState)
			}
		}
	}
}

func (cb *CircuitBreaker) nextState(state State) {
	cb.Mutex.Lock()
	defer cb.Mutex.Unlock()
	cb.Counts.Reset()
	switch cb.State {
	case CloseState:
		cb.State = OpenState
		cb.expireTime = time.Now().Add(cb.TimeOut)
		break
	case OpenState:
		cb.State = HalfOpenState
		break
	case HalfOpenState:
		if state == OpenState {
			cb.State = OpenState
			cb.expireTime = time.Now().Add(cb.TimeOut)
		} else {
			cb.State = CloseState
		}
	}
}

func DefaultToOpen(counts Counts) bool {
	return counts.TotalFailures > 10
}

func DefaultToClose(counts Counts) bool {
	return counts.TotalSuccesses > 10
}
