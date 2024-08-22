package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

type Employee struct {
	Person
	Position string
}

func (e Employee) Describe() {

	fmt.Println("Name: ", e.Name, "Last Name: ", e.LastName, "Age: ", e.Age, "Position: ", e.Position)
}

func main() {

	person1 := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
	}

	Employee1 := Employee{
		Person:   person1,
		Position: "Manager",
	}

	Employee1.Describe()
}
