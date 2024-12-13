package ratelimit

type Limiter interface {
	Allow(key string) bool
	GetCurrentCount(key string) uint64
}
