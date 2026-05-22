package main

import "fmt"

type User struct {
	Name string
}

var u = User{Name: "Syed"}

func update(user User) { // Call By Value
	user.Name = "Jafer"
}

func main() {

	update(u)

	fmt.Println(u.Name)
}
