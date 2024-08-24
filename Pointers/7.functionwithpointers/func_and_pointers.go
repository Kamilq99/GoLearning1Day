package main

import "fmt"

func changeVariableValue(a *int) {
	*a = 10
}

func main() {

	b := 5

	fmt.Println("Value of variable b is: ", b)

	changeVariableValue(&b)

	fmt.Println("After chaning value by pointer b is: ", b)

}
