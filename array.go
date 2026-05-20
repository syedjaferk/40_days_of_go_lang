package main

import "fmt"

func main() {
    // Pattern: var name [size]type
    var grades [5]int 
    grades[0] = 95 // Assigning a value

    colors := [...]string{"Red", "Green", "Blue"}

    fmt.Println(colors);
    fmt.Println(grades);

    arr := [3]string{"Go", "Python", "Rust"}

    for i := 0; i < len(arr); i++ {
        fmt.Printf("Index %d holds value %s\n", i, arr[i])
    }

    primes := [5]int{2, 3, 5, 7, 11}

    for index, value := range primes {
        fmt.Println(index, value)
    }


}