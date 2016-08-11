package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pt := Point{1, 2}
	fmt.Println(pt.X)
	fmt.Println(pt.Y)
}
