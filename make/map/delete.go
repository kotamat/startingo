package main

import "fmt"

func main() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	delete(m, 2)
	fmt.Println(m)
}
