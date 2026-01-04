package ratelimit

import (
	"sync"
	"time"
)

type Limiter interface {
	Allow(key string) bool
}

type TokenBucketLimit struct {
	buckets map[string]*bucket
	mu      sync.Mutex

	capacity   int
	refillRate float64
	ttl        time.Duration
}

func NewLimiter() Limiter {
	return NewTokenBucket()
}

func NewTokenBucketLimiter(capacity int, refillRate float64, ttl time.Duration) *TokenBucketLimit {
	limiter := &TokenBucketLimit{
		buckets:    make(map[string]*bucket),
		capacity:   capacity,
		refillRate: refillRate,
		ttl:        ttl,
	}

	go limiter.cleanupLoop()

	return limiter
}

func (l *TokenBucketLimit) Allow(key string) bool {
	l.mu.Lock()
	b, exists := l.buckets[key]

	if !exists {
		b = newBucket(l.capacity, l.refillRate)
		l.buckets[key] = b
	}

	l.mu.Unlock()

	return b.allow()
}
