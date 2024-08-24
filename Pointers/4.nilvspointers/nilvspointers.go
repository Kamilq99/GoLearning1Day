package main

import (
	"fmt"
)

func main() {
	var value *int = nil

	if value == nil {
		fmt.Println("Value is nil")
	} else {
		fmt.Println("Value is not nil")
	}
}
