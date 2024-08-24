package main

import (
	"fmt"
)

func main() {
	ptr := new(int)

	*ptr = 42

	fmt.Println("Variable Value:", *ptr)
}
