package main

import "fmt"

func main() {
	names := make([]string, 0, 10)
	// var names []string
	fmt.Println("Names ", names)

	names = append(names, "Syed", "Wijay")

	fmt.Println(names)

	for _, value := range names {
		fmt.Printf("Value %s\n", value)
	}
}