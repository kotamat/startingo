package main

import "fmt"

func main() {
	as()

	asp()
}

func as() {
	// 配列からのスライスを変更すると、元の配列も変更される
	// 逆も然り
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:2]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	a[1] = 0
	s[0] = 9
	fmt.Println(a) // => [9 0 3 4 5]
	fmt.Println(s) // => [9 0]

}

func asp() {
	// スライスの領域を拡張すると、もとの配列に対する参照は消える
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	a[0] = 9
	fmt.Println(a) // => [9 2 3]
	fmt.Println(s) // => [1 2 3 4]

}
