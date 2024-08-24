package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	fmt.Println("\n")

	slice = append(slice, 6, 7, 8, 9, 10)

	fmt.Println(slice)
}
