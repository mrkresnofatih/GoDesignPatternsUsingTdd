package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// enrichContextStudyCase()
	contextWithTimeoutStudyCase()
}

// Enrich Context Data - Study Case
func enrichContextStudyCase() {
	fmt.Println("enrich ctx study case!")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value1")
	ctx = context.WithValue(ctx, "key2", "value2")
	fmt.Println(ctx.Value("key1"))
	fmt.Println(ctx.Value("key2"))
}

// Context with Timeout - Study Case
func contextWithTimeoutStudyCase() {
	fmt.Println("context deadline study case!")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var task sync.WaitGroup
	task.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timed out")
				task.Done()
				return
			default:
				fmt.Println("doing something cool")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	doneSignal := <-ctx.Done()
	fmt.Println(doneSignal)
	task.Wait()

	// OUTPUT
	// ======
	// context deadline study case!
	// doing something cool
	// doing something cool
	// doing something cool
	// doing something cool
	// doing something cool
	// doing something cool
	// {}
	// timed out

	// EXPLANATION
	// ===========
	// starts out with one main thread
	// creates a context that will terminate after 3s
	// creates a separate thread ("new thread")
	// at the same time, main thread continues & new thread executes
	// the ctx.Done() will return a channel at termination that is being listened by "main thread" & "new thread"
	// main thread receives channel and prints
	// new thread receives channel & prints "timed out" & returns
}
