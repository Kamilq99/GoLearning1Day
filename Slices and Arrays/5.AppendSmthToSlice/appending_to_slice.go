package main

import "fmt"

func AppendToSlice(slice []int, value int) []int {
	slice = append(slice, value)
	return slice
}

func main() {

	slice := []int{1, 2, 3, 4, 5}

	value := 6

	fmt.Println(AppendToSlice(slice, value))

}
