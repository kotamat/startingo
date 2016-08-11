package main

import "fmt"

func main() {
	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Printf("n(%d) is even \n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}
}
