package main

import "fmt"

func main() {

	change := func(x *int) {
		*x = 10
	}

	a := 5
	change(&a)
	fmt.Println(a)
}
