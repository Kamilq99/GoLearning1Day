package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	var arr2 [5]int = arr

	arr2[0] = 10

	fmt.Println("Values of arr is: ", arr)
	fmt.Println("\n")
	fmt.Println("Values of arr2 is: ", arr2)
}
