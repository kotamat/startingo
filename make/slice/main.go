package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:2]               //変数sは[]int型
	fmt.Println(s)            // => [1, 2]
	fmt.Println(a[len(a)-2:]) // [4, 5]
}
