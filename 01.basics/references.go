package main

import "fmt"

func main(){

	m1 := map[string]int {
		"a": 1,
	}

	m2 := m1

	m2["a"] = 100

	fmt.Println(m1)
	fmt.Println(m2)

}