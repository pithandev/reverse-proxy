package ratelimit

type Limiter interface {
	Allow(key string) bool
}

func NewLimiter() Limiter {
	return NewTokenBucket()
}
