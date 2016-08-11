package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s = append(s, 4)
	fmt.Println(s)

	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)
	fmt.Println(s2)

	//copy
	c1 := []int{1, 2, 3, 4, 5}
	c2 := []int{10, 11}
	n := copy(c1, c2)
	fmt.Println(n, c1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a1 := a[2:6]
	fmt.Println(len(a1))
	fmt.Println(cap(a1))
	a2 := a[2:4:4]
	fmt.Println(len(a2))
	fmt.Println(cap(a2))
	a3 := a[2:4:6]
	fmt.Println(len(a3))
	fmt.Println(cap(a3))
}
