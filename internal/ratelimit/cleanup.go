package ratelimit

import "time"

func (l *TokenBucketLimit) cleanupLoop() {
	ticker := time.NewTicker(l.ttl)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		l.mu.Lock()
		for key, b := range l.buckets {
			b.mu.Lock()
			idle := now.Sub(b.lastRefill) > l.ttl
			b.mu.Unlock()

			if idle {
				delete(l.buckets, key)
			}
		}
		l.mu.Unlock()
	}
}
