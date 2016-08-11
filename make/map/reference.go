package main

import "fmt"

func main() {
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	var (
		s  string
		ok bool
	)

	s, ok = m[1]
	fmt.Println(s, ok)
	s, ok = m[9]
	fmt.Println(s, ok)
	_, ok = m[3]
	fmt.Println(s, ok)
}
