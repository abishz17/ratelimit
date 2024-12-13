# Go Sliding Window RateLimiter
A simple  rate limiting library for Go using the Sliding Window algorithm. Control the number of requests a user can make within a specified time window.

# Installation
```bash
go get github.com/abishz17/ratelimit
 ```
#Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/abishz17/ratelimit"
)

func main() {
	// Create a new rate limiter with a limit of 10 requests per 10 minutes
	limiter := ratelimit.NewRateLimiter(
		ratelimit.WithLimit(10),
		ratelimit.WithTimeWindow(10*time.Minute),
	)

	userID := "user1"

	// Simulate 20 requests
	for i := 0; i < 20; i++ {
		if limiter.Allow(userID) {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request Denied")
		}
	}
}
```
# API
```go
NewRateLimiter(opts ...Options) Limiter
Creates a new rate limiter with optional configurations.

Allow(key string) bool
Checks if a request is allowed for the given key. Returns true if allowed, false otherwise.

GetCurrentCount(key string) uint64
Retrieves the current count of requests for the given key.

Options
WithLimit(limit uint64): Sets the maximum number of requests allowed.
WithTimeWindow(timeWindow time.Duration): Sets the time window for rate limiting.
```


