package main

import "fmt"

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	// closure

	f2 := later()
	fmt.Println(f2("Golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome!"))
}
