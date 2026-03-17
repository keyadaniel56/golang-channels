# Worker Pool Pattern in Go

This project demonstrates the worker pool pattern using Go channels, where a master distributes tasks to multiple worker goroutines.

## Key Concepts

- **Fan-out**: Multiple goroutines reading from the same channel
- **Work Distribution**: Tasks are automatically balanced across workers
- **Result Collection**: Separate goroutine collects and processes results
- **Graceful Shutdown**: Proper channel closing and WaitGroup usage

## How It Works

1. Master goroutine generates tasks and sends them to a shared channel
2. Multiple worker goroutines compete for tasks from the channel
3. Each worker processes tasks and sends results to a results channel
4. Collector goroutine gathers and displays all results
5. WaitGroup ensures all workers complete before program exits

## Run the Example

```bash
go run main.go task.go