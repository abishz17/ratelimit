package ratelimit

import "time"

type Options func(*RateLimitOptions)

func WithLimit(limit uint64) Options {
	return func(o *RateLimitOptions) {
		o.Limit = limit
	}
}

func WithTimeWindow(timeWindow time.Duration) Options {
	return func(o *RateLimitOptions) {
		o.timeWindow = timeWindow
	}
}
