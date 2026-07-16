package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func increment() {

	defer wg.Done()

	for i := 0; i < 10000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
func main() {
	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println(counter)
}
