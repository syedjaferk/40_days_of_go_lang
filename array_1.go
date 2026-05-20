package main

import "fmt"

func main() {
	// var array [3]int 
	// array[0] = 10
	// array[1] = 20
	// array[2] = 30

	// fmt.Println(array, array[0])

	// array[2] = 234

	// fmt.Println(array)

	langs := [3]string{"Go", "Python", "Rust"}


	for itr := 0; itr < len(langs); itr ++ {
		fmt.Printf("Index %d - Value %s\n", itr, langs[itr])
	}

	fmt.Println("Using Range")

	for index, value := range langs {
		fmt.Printf("Index %d - Value %s\n", index, value)
	}

	for _, value := range langs {
		fmt.Printf("Value %s\n", value)
	}

}