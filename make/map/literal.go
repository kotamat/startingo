package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	fmt.Println(m)

	// slice in map
	ms := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(ms)
}
