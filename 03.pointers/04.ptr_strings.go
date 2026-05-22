package main

import "fmt"

func main() {
	name := "Syed Jafer"

	var address *string = &name

	fmt.Println("Name ", name)
	fmt.Println("Address ", address)
	fmt.Println("Value from Pointer ", *address)

	*address = "Arun"

	fmt.Println("Value from Pointer ", *address)
	fmt.Println("Name ", name)

}
