package main

import "fmt"

func main() {
	a := [3]string{"Apple", "Banana", "Cherry"}
	p := &a
	fmt.Println(a[1])
	fmt.Println(p[1]) // (*p)[1] に自動変換してくれる
	p[2] = "Grape"    // (*p)[2] に自動変換してくれる
	fmt.Println(a[2])
	fmt.Println(p[2])

	// アドレス
	fmt.Printf("type=%T, address=%p\n", p, p)
}
