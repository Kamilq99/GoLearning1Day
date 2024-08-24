package main

import "fmt"

func changeValue(value *int) {
	*value = 10
}

func main() {

	a := 5

	fmt.Println("Value before changes: ", a)

	changeValue(&a)

	fmt.Println("Value after changes: ", a)
}
