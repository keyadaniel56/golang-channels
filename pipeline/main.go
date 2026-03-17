package main

import "fmt"

func main() {
    fmt.Println("Starting Pipeline Pattern Example")
    fmt.Println("=================================")
    
    // Build the pipeline
    // Stage 1: Generate data
    dataStream := generator(10)
    
    // Stage 2: Transform data with 3 parallel workers (fan-out)
    transformedStream := transformer(dataStream, 3)
    
    // Stage 3: Save/display results
    saver(transformedStream)
    
    fmt.Println("\nPipeline completed!")
}