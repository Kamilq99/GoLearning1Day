package main

import "fmt"

func Array_vs_Slice(arr []int, slice []int) {

	fmt.Println("Original Values of arr is: ", arr)
	fmt.Println("Original Values of slice is: ", slice)

	fmt.Println("\n")

	arr[3] = 7
	slice[3] = 7

	fmt.Println("Changed Values of arr is: ", arr)
	fmt.Println("Changed Values of slice is: ", slice)

}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	slice := []int{1, 2, 3, 4, 5}

	Array_vs_Slice(arr, slice)

}
