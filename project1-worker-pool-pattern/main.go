package main

import (
	"fmt"

	"sync"
	"time"
)

/**
 * Handles main
 */
func main() {
	// Create channels for tasks and results
	taskCh := make(chan Task, 10)     // Buffered channel for tasks
	resultCh := make(chan Result, 10) // Buffered channel for results

	// Configuration
	numWorkers := 3
	numTasks := 20

	// WaitGroup to track workers
	var wg sync.WaitGroup

	// Start workers (fan-out)
	fmt.Printf("Starting %d workers...\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		/**
		 * Handles worker
		 */
		go worker(i, taskCh, resultCh, &wg)
	}

	// Start result collector
	done := make(chan bool)
	/**
	 * Handles collector
	 */
	go collector(resultCh, done)

	// Master: Distribute tasks (fan-out)
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := NewTask(i)
			fmt.Printf("Master: Sending %s to workers\n", task.Data)
			taskCh <- task
			time.Sleep(50 * time.Millisecond) // Simulate task generation delay
		}
		/**
		 * Handles close
		 */
		close(taskCh) // No more tasks to send
		fmt.Println("Master: All tasks sent")
	}()

	// Wait for all workers to finish
	wg.Wait()
	/**
	 * Handles close
	 */
	close(resultCh) // Signal collector that no more results are coming

	// Wait for collector to finish
	<-done

	fmt.Println("\nAll tasks completed successfully!")
}

// worker processes tasks from the task channel
func worker(id int, taskCh <-chan Task, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskCh { // Keep processing until channel is closed
		fmt.Printf("Worker %d: Started %s\n", id, task.Data)
		result := task.Process(id)
		resultCh <- result
		fmt.Printf("Worker %d: Completed %s\n", id, task.Data)
	}

	fmt.Printf("Worker %d: Shutting down\n", id)
}

// collector gathers and displays results
func collector(resultCh <-chan Result, done chan<- bool) {
	results := []Result{}

	for result := range resultCh { // Keep collecting until channel is closed
		fmt.Printf("Collector: Received result: Task %d processed by Worker %d - %s\n",
			result.TaskID, result.WorkerID, result.Output)
		results = append(results, result)
	}

	fmt.Printf("\nSummary: Collected %d results\n", len(results))
	done <- true
}
