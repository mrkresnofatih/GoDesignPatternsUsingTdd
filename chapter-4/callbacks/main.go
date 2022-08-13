package main

import (
	"fmt"
	"sync"
)

func main() {
	// executeSyncStudyCase()

	executeAsyncStudyCase()
}

// sync
func executeSyncStudyCase() {
	GetSequenceValueByOrder(7, func(value int) {
		fmt.Printf("sequence value is %d\n", value)
	})
}

// async
func executeAsyncStudyCase() {
	var task sync.WaitGroup
	task.Add(1)
	go GetSequenceValueByOrder(7, func(value int) {
		fmt.Printf("sequence value is %d\n", value)
		task.Done()
	})
	task.Wait()
}

func GetSequenceValueByOrder(order int, callback func(int)) {
	callback(3 + (order-1)*4)
}
