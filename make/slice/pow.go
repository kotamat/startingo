package main

import "fmt"

func pow(s []int) {
	for i, v := range s {
		s[i] = v * v

	}
	return
}

func main() {
	s := []int{1, 2, 3}
	pow(s)
	fmt.Println(s)
}
