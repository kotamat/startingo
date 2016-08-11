package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	ps := make([]Point, 5)
	for _, p := range ps {
		fmt.Println(p)

	}
}
