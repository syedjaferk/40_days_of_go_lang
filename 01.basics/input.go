package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// var name string;
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name ")
	name, _ := reader.ReadString('\n')
	// fmt.Scanf("%s", &name)


	var a, b string 
	fmt.Println("Enter values of a and b")

	fmt.Scanf("%s %s", &a, &b)


	fmt.Println("Hello ", name, a, b)
}