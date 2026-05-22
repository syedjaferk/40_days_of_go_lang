package main

import "fmt"

type User struct {
	Name string
}

var u = User{Name: "Syed"}

func update(user *User) { // Call By Reference
	(*user).Name = "Jafer" // Go Struct -> default de reference.
	// user.Name -> (*user).Name
}

func main() {

	update(&u)

	fmt.Println(u.Name)
}
