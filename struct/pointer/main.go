package main

import "fmt"

func main() {
	var i int

	i = 5
	p := &i
	fmt.Printf("%T\n", p)
	pp := &p
	fmt.Printf("%T\n", pp)

	*p = 10
	fmt.Println(i)
}
