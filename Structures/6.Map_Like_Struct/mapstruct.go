package main

import "fmt"

type Person struct {
	Name       string
	LastName   string
	FubtolClub string
}

func main() {
	futboller_1 := Person{
		Name:       "Cristiano",
		LastName:   "Ronaldo",
		FubtolClub: "Al-Nassr",
	}

	futboller_2 := Person{
		Name:       "Lionel",
		LastName:   "Messi",
		FubtolClub: "Inter Miami",
	}

	futboller_3 := Person{
		Name:       "Neymar",
		LastName:   "Jr",
		FubtolClub: "Al-Hilal",
	}

	futboller_4 := Person{
		Name:       "Mohamed",
		LastName:   "Salah",
		FubtolClub: "Liverpool FC",
	}

	m := map[string]Person{
		"Cristiano": futboller_1,
		"Lionel":    futboller_2,
		"Neymar":    futboller_3,
		"Mohamed":   futboller_4,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
