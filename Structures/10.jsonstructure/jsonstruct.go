package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      int    `json:"age"`
}

func main() {

	person := Person{
		Name:     "David",
		LastName: "Beckham",
		Age:      40,
	}

	jsonPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(jsonPerson))

	var deserializedPerson Person

	err = json.Unmarshal(jsonPerson, &deserializedPerson)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(deserializedPerson)
}
