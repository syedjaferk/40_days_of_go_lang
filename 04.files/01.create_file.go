package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")

	// To check an error
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Println("File Created Successfully")
}
