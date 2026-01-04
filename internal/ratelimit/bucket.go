package ratelimit

import (
	"sync"
	"time"
)

type tokenBucket struct{}

type bucket struct {
	capacity   int
	tokens     float64
	refillRate float64
	lastRefill time.Time
	mu         sync.Mutex
}

func newBucket(capacity int, refillRate float64) *bucket {
	return &bucket{
		capacity:   capacity,
		tokens:     float64(capacity),
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (b *bucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()

	b.lastRefill = now

	b.tokens += elapsed * b.refillRate
	if b.tokens > float64(b.capacity) {
		b.tokens = float64(b.capacity)
	}

	if b.tokens < 1 {
		return false
	}

	b.tokens--
	return true
}

func NewTokenBucket() Limiter {
	return &tokenBucket{}
}

func (t *tokenBucket) Allow(key string) bool {
	return true
}
