package main

import "fmt"

// Function to simulate pointer arithmetic for array search
func search(arr []int, target int) int {
	// We can use a pointer to simulate the behavior.
	// Since Go does not support pointer arithmetic, we use array indexing instead.
	for i := 0; i < len(arr); i++ {
		// Simulate dereferencing the pointer by accessing the array element
		if arr[i] == target {
			return i // Return the index of the target element
		}
	}
	return -1 // Return -1 if the target element is not found
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	target := 30

	// Call the search function
	index := search(arr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Println("Element not found")
	}
}
