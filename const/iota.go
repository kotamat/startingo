package main

import "fmt"

const (
	A = iota // 0
	B
	C
)

func main() {
	fmt.Printf("%v %v %v", A, B, C)
}
