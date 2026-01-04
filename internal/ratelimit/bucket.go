package ratelimit

type tokenBucket struct{}

func NewTokenBucket() Limiter {
	return &tokenBucket{}
}

func (t *tokenBucket) Allow(key string) bool {
	return true
}
