package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("\n Hi How are you. Welcome to leanring group")
}
