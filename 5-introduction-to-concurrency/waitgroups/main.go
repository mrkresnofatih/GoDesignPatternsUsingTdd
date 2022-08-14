package main

import (
	"fmt"
	"sync"
)

func main() {
	var task sync.WaitGroup
	numOfConcurrentProcesses := 10
	task.Add(numOfConcurrentProcesses)
	for index := 0; index < numOfConcurrentProcesses; index++ {
		go func(number int) {
			fmt.Printf("Hello world by process no.: %d\n", number)
			task.Done()
		}(index)
	}
	task.Wait()
}
