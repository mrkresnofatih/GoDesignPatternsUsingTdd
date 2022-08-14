package main

import (
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	var task sync.WaitGroup
	counter := Counter{}
	for i := 0; i < 10; i++ {
		task.Add(1)
		go func(i int) {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
			task.Done()
		}(i)
	}
	task.Wait()
	counter.Lock()
	defer counter.Unlock()
	println(counter.value)
}
