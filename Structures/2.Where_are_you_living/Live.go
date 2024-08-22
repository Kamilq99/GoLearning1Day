package main

import "fmt"

type Living struct {
	Address string
	City    string
	Postal  string
}

type Person struct {
	Name        string
	LastName    string
	Age         int
	LivingPlace Living
}

func defCredentials() Person {

	credentials := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
		LivingPlace: Living{
			Address: "123 Main St.",
			City:    "Miami",
			Postal:  "12345",
		},
	}

	fmt.Println("Name: ", credentials.Name)
	fmt.Println("Last Name: ", credentials.LastName)
	fmt.Println("Age: ", credentials.Age)
	fmt.Println("Living: ", credentials.LivingPlace)

	return credentials
}

func main() {

	defCredentials()
}
