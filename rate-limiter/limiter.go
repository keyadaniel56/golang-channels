

package main

import (
    "fmt"
    "time"
)

// Request represents an API request
type Request struct {
    ID        int
    Endpoint  string
    Timestamp time.Time
}

// RateLimiter controls the rate of request processing
type RateLimiter struct {
    ticker     *time.Ticker
    requests   chan Request
    quit       chan bool
    rate       int           // requests per second
    processed  int
    limited    int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate int) *RateLimiter {
    return &RateLimiter{
        ticker:   time.NewTicker(time.Second / time.Duration(rate)),
        requests: make(chan Request, 100),
        quit:     make(chan bool),
        rate:     rate,
    }
}

// Start begins processing requests at the limited rate
func (rl *RateLimiter) Start() {
    fmt.Printf("Rate limiter started: %d requests/second\n", rl.rate)
    
    for {
        select {
        case req := <-rl.requests:
            // Process the request
            rl.processed++
            fmt.Printf("[%s] Processing request %d: %s (Processed: %d, Limited: %d)\n",
                time.Now().Format("15:04:05.000"),
                req.ID, req.Endpoint, rl.processed, rl.limited)
            
        case <-rl.ticker.C:
            // Rate limiter tick - ready for next request
            continue
            
        case <-rl.quit:
            fmt.Println("Rate limiter stopping...")
            rl.ticker.Stop()
            return
        }
    }
}

// Submit tries to submit a request to the rate limiter
func (rl *RateLimiter) Submit(req Request) bool {
    select {
    case rl.requests <- req:
        return true // Request accepted
    default:
        // Channel is full - request would exceed rate limit
        rl.limited++
        fmt.Printf("[%s] ⚠️  Request %d limited (rate exceeded)\n",
            time.Now().Format("15:04:05.000"), req.ID)
        return false
    }
}

// Stop shuts down the rate limiter
func (rl *RateLimiter) Stop() {
    rl.quit <- true
    close(rl.requests)
}