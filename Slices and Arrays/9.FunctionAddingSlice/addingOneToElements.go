package main

import "fmt"

func addingOneToSliceElements(slice []int) {

	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] + 1
	}

	fmt.Println(slice)

}

func main() {

	slice := []int{1, 2, 3, 4, 5}

	addingOneToSliceElements(slice)
}
