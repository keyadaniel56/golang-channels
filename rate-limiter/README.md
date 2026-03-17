# Rate Limiter Pattern in Go

This project demonstrates rate limiting using Go's `time.Ticker` and channels to control request processing rates.

## Key Concepts

- **time.Ticker**: Provides periodic ticks for rate control
- **select Statement**: Handles multiple channel operations
- **Buffered Channels**: Acts as a request queue
- **Non-blocking Operations**: Using select with default case

## How It Works

1. **Rate Limiter**: Controls processing to N requests per second
2. **Request Submission**: Attempts to queue requests
3. **Tick-based Processing**: Processes one request per tick
4. **Overflow Handling**: Rejects requests when queue is full

## Rate Limiting Strategies Demonstrated

- **Burst Handling**: Shows how rate limiter handles sudden spikes
- **Steady State**: Processing at consistent rate
- **Request Queue**: Buffered channel acts as temporary storage
- **Rejection**: Requests beyond capacity are immediately rejected

## Run the Example

```bash
go run main.go limiter.go

or

go run *.go