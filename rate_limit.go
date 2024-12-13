package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

type window struct {
	count     uint64
	timestamp time.Time
}

type SlidingWindow struct {
	previous *window
	current  *window
}

type RateLimitOptions struct {
	Limit      uint64
	timeWindow time.Duration
}

type RateLimiter struct {
	sync.RWMutex
	maxRequests uint64
	windows     map[string]*SlidingWindow
	options     RateLimitOptions
}

func (rl *RateLimiter) SetLimit(limit uint64) {
	rl.options.Limit = limit
}

func (rl *RateLimiter) SetTimeWindow(timeWindow time.Duration) {
	rl.options.timeWindow = timeWindow
}

func NewRateLimiter(opts ...Options) Limiter {
	rlOpts := &RateLimitOptions{
		Limit:      20,
		timeWindow: 60 * time.Second,
	}
	for _, opt := range opts {
		opt(rlOpts)
	}

	rl := &RateLimiter{
		windows: make(map[string]*SlidingWindow),
		options: *rlOpts,
	}
	return rl
}

func (r *RateLimiter) Allow(key string) bool {
	r.Lock()
	defer r.Unlock()
	now := time.Now()
	slwindow, exists := r.windows[key]
	if !exists {
		slwindow = &SlidingWindow{
			previous: &window{count: 0, timestamp: now},
			current:  &window{count: 0, timestamp: now},
		}
		r.windows[key] = slwindow
	}
	// If current window expires the current window becomes previous and new current window is initiated.
	timeElapsed := now.Sub(slwindow.current.timestamp)
	if timeElapsed >= r.options.timeWindow {
		slwindow.previous = slwindow.current
		slwindow.current = &window{count: 0, timestamp: now}
	}

	weight := 1 - (float64(timeElapsed) / float64(r.options.timeWindow))
	fmt.Println("Weight:", weight)
	if weight < 0 {
		weight = 0
	}

	// Calculate total count including weighted previous window
	weightedPreviousCount := uint64(float64(slwindow.previous.count) * weight)
	totalCount := weightedPreviousCount + slwindow.current.count
	if totalCount >= r.options.Limit {
		return false
	}
	slwindow.current.count++
	return true
}

func (r *RateLimiter) GetCurrentCount(key string) uint64 {
	r.RLock()
	defer r.RUnlock()
	slwindow, exists := r.windows[key]
	if !exists {
		return 0
	}
	return slwindow.current.count
}
