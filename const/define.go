package main

import "fmt"

const (
	X = 1
	Y
	Z
	S1 = "あ"
	S2
	S = S1 + S2
)

func main() {
	fmt.Printf("%v %v %v %v %v %v", X, Y, Z, S1, S2, S)
}
