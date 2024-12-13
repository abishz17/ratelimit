package main

import (
	"fmt"
	ratelimit "ratelimiter"
	"time"
)

func main() {
	limiter := ratelimit.NewRateLimiter(
		ratelimit.WithLimit(10),
		ratelimit.WithTimeWindow(10*time.Minute),
	)
	userId := "user1"
	for i := 0; i < 20; i++ {
		if limiter.Allow(userId) {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request Denied")
		}
	}
}
