package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

type Person struct {
	Name     string
	Age      int
	IsActive bool
	Salary   float64
	Address  Address
}

func (p Person) IsEligibleToVote() bool {
	if p.Age < 18 {
		return false
	}
	return true
}

func main() {
	person := Person{
		Name: "Jafer",
		Age:  25,
		Address: Address{
			Street: "2nd Street Chitra Nagar",
			City:   "Coimbatore",
			State:  "Tamil Nadu",
		},
	}

	fmt.Println(person)

	fmt.Println("Name ", person.Name)
	fmt.Println("Age ", person.Age)

	person.Age = 28

	fmt.Println(person.IsEligibleToVote())
}
