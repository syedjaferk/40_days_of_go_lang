package main

import "fmt"

func main(){

	ages := make(map[string]int)

	ages["Vignesh"] = 26
	ages["Jafer"] = 28
	ages["Syed"] = 28
	ages["Aron"] = 28
	// ages["baby"] = 0

	// value, exists := ages["baby"]

	// fmt.Println(value, exists)

	fmt.Println(ages["Food"])


	
}