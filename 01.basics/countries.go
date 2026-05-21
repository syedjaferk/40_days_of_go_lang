package main

import "fmt"

func main(){

	countries := map[string]string {
		"IN": "India",
		"LK": "SriLanka",
		"JP": "Japan",
	}

	for key, value := range countries {
		fmt.Println(key, " ====> " , value)
	}

	countries["US"] = "United States"

	for key, value := range countries {
		fmt.Println(key, " ====> " , value)
	}

	fmt.Println(len(countries))

}