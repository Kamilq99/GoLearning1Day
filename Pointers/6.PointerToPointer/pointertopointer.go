package main

import (
	"fmt"
)

func main() {

	a := 5

	fmt.Println(a)

	b := &a

	fmt.Println(*b)

	c := &b

	**c = 10

	fmt.Println(**c)
}
