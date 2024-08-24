package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice

	fmt.Println("Original slice: ", slice)
	fmt.Println("Copied slice: ", slice2)

	slice2[0] = 248
	fmt.Println("Copied slice with changing element: ", slice2)
}
