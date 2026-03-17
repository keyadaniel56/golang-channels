package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Rate Limiter Pattern Example")
    fmt.Println("============================")
    
    // Create a rate limiter allowing 3 requests per second
    limiter := NewRateLimiter(3)
    
    // Start the rate limiter in a goroutine
    go limiter.Start()
    
    // Simulate incoming requests at varying rates
    endpoints := []string{"/api/users", "/api/products", "/api/orders", "/api/payments"}
    
    fmt.Println("\nSimulating burst of requests...")
    
    // Send 15 requests in a burst
    for i := 1; i <= 15; i++ {
        req := Request{
            ID:        i,
            Endpoint:  endpoints[i%len(endpoints)],
            Timestamp: time.Now(),
        }
        
        accepted := limiter.Submit(req)
        
        if accepted {
            fmt.Printf("✓ Request %d submitted to %s\n", i, req.Endpoint)
        }
        
        // Small delay between submissions
        time.Sleep(50 * time.Millisecond)
    }
    
    fmt.Println("\nSimulating steady stream...")
    
    // Send requests at a steady rate (slightly above limit)
    for i := 16; i <= 30; i++ {
        req := Request{
            ID:        i,
            Endpoint:  endpoints[i%len(endpoints)],
            Timestamp: time.Now(),
        }
        
        limiter.Submit(req)
        
        // 400ms between requests = 2.5 req/sec (slightly under 3/sec limit)
        time.Sleep(400 * time.Millisecond)
    }
    
    // Let the limiter process remaining requests
    fmt.Println("\nWaiting for remaining requests to process...")
    time.Sleep(2 * time.Second)
    
    // Stop the rate limiter
    limiter.Stop()
    
    fmt.Printf("\nFinal Stats:\n")
    fmt.Printf("  Processed: %d requests\n", limiter.processed)
    fmt.Printf("  Limited: %d requests\n", limiter.limited)
    fmt.Printf("  Total submitted: %d\n", limiter.processed+limiter.limited)
}