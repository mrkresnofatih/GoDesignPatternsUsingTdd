package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channelCaseStudy()
	channelWithNewThreadLeadCaseStudy()
	// channleWithMainThreadLeadCaseStudy()
}

func channelCaseStudy() {
	channel := make(chan int)
	go func() {
		channel <- 3
	}()

	val := <-channel
	fmt.Printf("Value: %d\n", val)
}

func channelWithNewThreadLeadCaseStudy() {
	fmt.Println("start")
	channel := make(chan int)
	var task sync.WaitGroup
	task.Add(1)
	go func() {
		fmt.Println("Finishing the goroutine")
		channel <- 4
		task.Done()
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("Main thread wait finished")
	value := <-channel
	fmt.Println(value)
	task.Wait()

	// OUTPUT
	// ======
	// start
	// Finishing the goroutine
	// Main thread wait finished
	// 4

	// EXPLANATION
	// ===========
	// starts out with 1 thread
	// sees a go routine, creates a separate thread
	// at the same time, "main thread" continues & "new thread" executes
	// execution of "new thread" finishes earlier, returns the channel a value
	// main thread has 3s difference with new thread, has to wait
	// main thread prints the submitted value and main thread terminates
}

func channleWithMainThreadLeadCaseStudy() {
	fmt.Println("start")
	channel := make(chan int)
	var task sync.WaitGroup
	task.Add(1)
	go func() {
		fmt.Println("Finishing the goroutine")
		time.Sleep(time.Second * 3)
		fmt.Println("New thread wait finished")
		channel <- 8
		task.Done()
	}()

	value := <-channel
	fmt.Println(value)
	task.Wait()

	// OUTPUT
	// ======
	// start
	// Finishing the goroutine
	// New thread wait finished
	// 8

	// EXPLANATION
	// ===========
	// starts out with 1 thread
	// sees a go routine, creates a separate thread
	// at the same time, "main thread" continues & "new thread" executes
	// execution of "main thread" finishes earlier & channel hasn't got a value submitted yet
	// main thread has to wait for new thread to submit a value to the channel to continue
	// when value is submitted & received, main thread prints the submitted value
	// then main thread terminates
}
