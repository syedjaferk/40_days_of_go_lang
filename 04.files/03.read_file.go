package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("notes.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
