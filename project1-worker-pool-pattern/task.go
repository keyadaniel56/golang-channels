package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Task represents a unit of work to be processed
type Task struct {
    ID      int
    Data    string
    Duration time.Duration
}

// NewTask creates a new task with random processing time
func NewTask(id int) Task {
    durations := []time.Duration{
        100 * time.Millisecond,
        250 * time.Millisecond,
        500 * time.Millisecond,
        1 * time.Second,
    }
    return Task{
        ID:      id,
        Data:    fmt.Sprintf("Task-%d", id),
        Duration: durations[rand.Intn(len(durations))],
    }
}

// Process simulates work being done on a task
func (t Task) Process(workerID int) Result {
    time.Sleep(t.Duration)
    return Result{
        TaskID:   t.ID,
        WorkerID: workerID,
        Output:   fmt.Sprintf("Processed %s in %v", t.Data, t.Duration),
    }
}

// Result represents the output of processing a task
type Result struct {
    TaskID   int
    WorkerID int
    Output   string
}