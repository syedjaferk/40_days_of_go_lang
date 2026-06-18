package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "Name of User")
	age := flag.Int("age", 0, "User Age")

	// flag.Parse()

	fmt.Println("Hello", *name, " Age ", *age)
}
