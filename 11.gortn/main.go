package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 1 : ", i)
		time.Sleep(time.Second)
	}
}

func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 2 : ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	task1()
	task2()

	time.Sleep(3 * time.Second)
}
