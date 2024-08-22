package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
	City     string
}

func newPerson(name string, lastname string, age int, city string) Person {
	return Person{
		Name:     name,
		LastName: lastname,
		Age:      age,
		City:     city,
	}
}

func main() {
	person := newPerson("Jan", "Kowal", 25, "Wroclaw")
	fmt.Println(person)

	var name string
	var lastname string
	var age int
	var city string

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Last Name: ")
	fmt.Scan(&lastname)
	fmt.Print("Age: ")
	fmt.Scan(&age)
	fmt.Print("City: ")
	fmt.Scan(&city)

	person = newPerson(name, lastname, age, city)
	fmt.Println(person)
}
