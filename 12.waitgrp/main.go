package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(name, "Started")
	time.Sleep(2 * time.Second)
	fmt.Println(name, "Finished")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go worker("Wijay", &wg)
	go worker("Arun", &wg)
	go worker("Vignesh", &wg)

	wg.Wait()
	fmt.Println("All workers completed")
}
