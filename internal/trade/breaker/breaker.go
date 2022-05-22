// Package breaker stops trade worker if it violates its failure threshold.
package breaker // import cb "github.com/elkopass/BITA/internal/trade/breaker"

import (
	"github.com/elkopass/BITA/internal/config"
	"time"
)

type CircuitBreaker struct {
	failuresTotal   *int
	lastFailureTime *time.Time
}

func NewCircuitBreaker() *CircuitBreaker {
	zero := 0
	now := time.Now()

	return &CircuitBreaker{
		failuresTotal:   &zero,
		lastFailureTime: &now,
	}
}

func (cb *CircuitBreaker) IncFailures() {
	*cb.failuresTotal++
	cb.updateState()
}

// WorkerMustExit returns true if trade worker is unhealthy and must be killed.
func (cb CircuitBreaker) WorkerMustExit() bool {
	return *cb.failuresTotal > config.CircuitBreakerMaxFailures
}

// updateState will "forgot" old failures if they exist.
func (cb *CircuitBreaker) updateState() {
	now := time.Now()
	prev := cb.lastFailureTime.Add(config.CircuitBreakerRefreshTime)

	for ; prev.Before(now); prev = prev.Add(config.CircuitBreakerRefreshTime) {
		if *cb.failuresTotal <= 0 {
			break
		}

		*cb.failuresTotal--
	}

	cb.lastFailureTime = &now
}
