package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
	Living   string
	Height   float64
	Weight   float64
}

func defPersonStiff() Person {
	person1 := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
		Living:   "Miami",
		Height:   1.78,
		Weight:   75.5,
	}

	fmt.Println("Name: ", person1.Name)
	fmt.Println("Last Name: ", person1.LastName)
	fmt.Println("Age: ", person1.Age)
	fmt.Println("Living: ", person1.Living)
	fmt.Println("Height: ", person1.Height)
	fmt.Println("Weight: ", person1.Weight)

	return person1
}

func defPersonbyUser(Name string, LastName string, Age int, Living string, Height float64, Weight float64) Person {

	person2 := Person{
		Name:     Name,
		LastName: LastName,
		Age:      Age,
		Living:   Living,
		Height:   Height,
		Weight:   Weight,
	}

	return person2
}

func main() {

	fmt.Println("First Person is:\n")

	defPersonStiff()

	fmt.Println("Second Person is:\n")

	var Name string
	var LastName string
	var Age int
	var Living string
	var Height float64
	var Weight float64

	fmt.Print("Name: ")
	fmt.Scan(&Name)

	fmt.Print("Last Name: ")
	fmt.Scan(&LastName)

	fmt.Print("Age: ")
	fmt.Scan(&Age)

	fmt.Print("Living: ")
	fmt.Scan(&Living)

	fmt.Print("Height: ")
	fmt.Scan(&Height)

	fmt.Print("Weight: ")
	fmt.Scan(&Weight)

	defPersonbyUser(Name, LastName, Age, Living, Height, Weight)

}
