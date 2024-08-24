package main

import "fmt"

func iterateInArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("Value of arr[", i, "] is:", arr[i])
	}

}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	iterateInArray(arr[:])

}
