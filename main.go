package main

import (
	"errors"
	"fmt"
	"time"
)

// CircuitBreaker struct
type CircuitBreaker struct {
	state             string
	failureCount      int
	failureThreshold  int
	resetTimeout      time.Duration
	lastAttemptedTime time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            "CLOSED",
		failureCount:     0,
		failureThreshold: threshold,
		resetTimeout:     timeout,
	}
}

func (cb *CircuitBreaker) Call(action func() error) error {
	if cb.state == "OPEN" && time.Since(cb.lastAttemptedTime) > cb.resetTimeout {
		cb.state = "HALF-OPEN"
	}
	if cb.state == "OPEN" {
		return errors.New("Circuit Breaker is in OPEN state, request denied")
	}

	err := action()
	if err != nil {
		cb.failureCount++
		if cb.failureCount > cb.failureThreshold {
			cb.state = "OPEN"
			cb.lastAttemptedTime = time.Now()
		}
		return err
	}

	cb.failureCount = 0
	cb.state = "CLOSED"
	return nil
}

func main() {
	cb := NewCircuitBreaker(2, 5*time.Second)
	var err error

	for i := 0; i < 10; i++ {
		err = cb.Call(func() error {
			// Simulating a failing function
			return errors.New("Service failure!")
		})

		if err != nil {
			fmt.Println("Error:", err)
		}

		time.Sleep(1 * time.Second)
	}

	time.Sleep(6 * time.Second)

	err = cb.Call(func() error {
		// Simulating a success function
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("State:", cb.state)

}
