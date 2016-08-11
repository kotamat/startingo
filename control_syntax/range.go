package main

import "fmt"

func main() {
	fruits := [3]string{"Apple", "Banana", "Cherry"}

	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// string

	for i, r := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

}
