package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (person Person) HaveBirthday(a int) {

	fmt.Println("Now he is not 17 years old")

	person.Age = person.Age + a

	fmt.Println("Now he is: ", person.Age)

}

func main() {

	var a int

	fmt.Println("If Victora is 17 years old? Please tell me am I wrong?: ")
	fmt.Scanln(&a)

	person := Person{
		Name:     "Victoria",
		LastName: "Beckham",
		Age:      17,
	}

	person.HaveBirthday(a)
}
