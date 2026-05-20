package main

import (
	"encoding/json"
	"fmt"
)

type Geo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Address struct {
	City string `json:"city"`
	Geo  Geo    `json:"geo"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func main() {
	data := `
		{
			"id": 1,
			"name": "Syed Jafer K",
			"email": "syedjaferk@parottasalna.com",
			"address":{
				"city": "Chennai",
				"geo":{
					"lat": 12.45,
					"lng": 80.56
				}
			}
		}
	`

	fmt.Println(data)

	var user User
	// Deserialization
	json.Unmarshal([]byte(data), &user)

	fmt.Println(user.Name)
	fmt.Println(user.Address.City)
	fmt.Println(user.Address.Geo.Lat)
}
