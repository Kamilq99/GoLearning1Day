package main

import "fmt"

type Person struct {
	Name     string
	LastName string
}

func compareStructs(person1 Person, person2 Person) bool {
	return person1 == person2
}

func main() {

	person1 := Person{
		Name:     "David",
		LastName: "Beckham",
	}

	var name string
	var lastname string

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Last Name: ")
	fmt.Scan(&lastname)

	person2 := Person{
		Name:     name,
		LastName: lastname,
	}

	fmt.Println(compareStructs(person1, person2))
}
