package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args)

	if len(os.Args) < 2 {
		fmt.Println("Usage app <name>")
		return
	}

	fmt.Println("Hello ", os.Args[1])
}
