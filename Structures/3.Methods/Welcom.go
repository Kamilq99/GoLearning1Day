package main

import "fmt"

type Person struct {
	Name     string
	LastName string
}

func (person Person) Greet() string {
	return fmt.Sprintf("Hello %s %s", person.Name, person.LastName)
}

func main() {

	person := Person{
		Name:     "David",
		LastName: "Beckham",
	}

	fmt.Println(person.Greet())
}
