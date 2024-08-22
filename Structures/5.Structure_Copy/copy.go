package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

func main() {

	var person1 Person
	var person2 Person

	person1.Name = "David"
	person1.LastName = "Beckham"
	person1.Age = 40

	fmt.Println("Name: ", person1.Name)
	fmt.Println("Last Name: ", person1.LastName)
	fmt.Println("Age: ", person1.Age)

	person2 = person1

	person2.Name = "Jan"
	person2.LastName = "Kowal"
	person2.Age = 25

	fmt.Println("Name: ", person2.Name)
	fmt.Println("Last Name: ", person2.LastName)
	fmt.Println("Age: ", person2.Age)
}
