package main

import "fmt"

func update(num *int) { // received a copy of x -> Call by Value.
	*num = 100
}

func main() {
	x := 10

	update(&x)

	fmt.Println(x)
}
