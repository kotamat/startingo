package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("%v", f(2, 3))
	fmt.Printf("%T", f)
}
