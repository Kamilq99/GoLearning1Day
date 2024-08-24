package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println("Value of arr[i] is: ", arr[i])
	}

	a := &arr[0]

	*a = 10

	for i := 0; i < len(arr); i++ {
		fmt.Println("Value of arr[i] is: ", arr[i])
	}

}
