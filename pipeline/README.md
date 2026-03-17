# Pipeline Pattern in Go

This project demonstrates the pipeline pattern using Go channels, showing how data flows through processing stages.

## Key Concepts

- **Channel Direction**: Functions return receive-only channels (`<-chan`)
- **Fan-out**: Multiple goroutines processing from one channel
- **Fan-in**: Multiple goroutines writing to one channel
- **Stage Composition**: Connecting channels to form a pipeline

## Pipeline Stages

1. **Generator**: Produces raw data items
2. **Transformer**: Converts data (uppercase) with parallel workers
3. **Saver**: Collects and displays final results

## Visual Pipeline


## Run the Example

```bash
go run main.go stages.go

or


## Run the Example

```bash
go run *.go