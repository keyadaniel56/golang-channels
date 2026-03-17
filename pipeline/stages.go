

package main

import (
    "fmt"
    "strings"
    "time"
)

// Data represents an item flowing through the pipeline
type Data struct {
    ID      int
    Content string
}

// Generator stage: produces data
func generator(numItems int) <-chan Data {
    out := make(chan Data)
    
    go func() {
        for i := 1; i <= numItems; i++ {
            out <- Data{
                ID:      i,
                Content: fmt.Sprintf("item_%d", i),
            }
            time.Sleep(100 * time.Millisecond) // Simulate work
        }
        close(out)
        fmt.Println("Generator: Finished")
    }()
    
    return out
}

// Transformer stage: processes data (fan-out and fan-in)
func transformer(input <-chan Data, numWorkers int) <-chan Data {
    out := make(chan Data)
    
    // Fan-out: Start multiple workers
    for w := 1; w <= numWorkers; w++ {
        go func(workerID int) {
            for data := range input {
                fmt.Printf("Transformer %d: Processing %s\n", workerID, data.Content)
                
                // Transform the data
                transformed := Data{
                    ID:      data.ID,
                    Content: strings.ToUpper(data.Content),
                }
                
                // Send to output channel (fan-in happens automatically
                // as all workers write to the same channel)
                out <- transformed
                
                time.Sleep(150 * time.Millisecond) // Simulate work
            }
            fmt.Printf("Transformer %d: Finished\n", workerID)
        }(w)
    }
    
    // Close output channel when all transformers are done
    go func() {
        // Wait for all transformers to finish
        // In a real implementation, you'd use a WaitGroup here
        // For simplicity, we'll just wait and then close
        time.Sleep(2 * time.Second)
        close(out)
    }()
    
    return out
}

// Saver stage: collects and displays results
func saver(input <-chan Data) {
    results := []Data{}
    
    for data := range input {
        fmt.Printf("Saver: Received %s (ID: %d)\n", data.Content, data.ID)
        results = append(results, data)
    }
    
    fmt.Printf("\nSaver: Collected %d items\n", len(results))
    for _, r := range results {
        fmt.Printf("  - ID %d: %s\n", r.ID, r.Content)
    }
}