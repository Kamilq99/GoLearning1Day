package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (person *Person) changeData() {

	person.Name = "Jan"
	person.LastName = "Kowal"
	person.Age = 25

	fmt.Println("Name: ", person.Name)
	fmt.Println("Last Name: ", person.LastName)
	fmt.Println("Now he is: ", person.Age)
}

func main() {

	person := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
	}

	fmt.Println("Name: ", person.Name)
	fmt.Println("Last Name: ", person.LastName)
	fmt.Println("Age:", person.Age)
	fmt.Println("\n")

	fmt.Println("After changing data using pointer\n")

	person.changeData()

}
