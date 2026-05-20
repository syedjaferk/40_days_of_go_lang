package main

import "fmt"

func main() {
	var age int = 1

	if age >= 18 {
		fmt.Println("Eligible to Vote")
	} else {
		fmt.Println("Not Eligible")
	}

	// for i := 1; i <= 10; i ++ {
	// 	fmt.Println(i)
	// }

	var itr int = 10

	// while look a like
	for itr >= 0 {
		fmt.Println(itr)
		itr--
	}

	// infinite loop
	// for {
	// 	fmt.Println("Runnnnn....")
	// }


}