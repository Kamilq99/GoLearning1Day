package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice[2:7]

	fmt.Println("Original slice: ", slice)
	fmt.Println("Slice from 3 to 7: ", slice2)
}
