package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (person *Person) increaseAge() {

	person.Age = person.Age + 1

	fmt.Printf("He had a Birthday and now have %d years", person.Age)
}

func main() {

	person := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
	}

	fmt.Println("How many years he had?")

	person.increaseAge()

}
