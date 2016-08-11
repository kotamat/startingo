package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("unit32 max value = %d\n", math.MaxUint32)

	s := `
  Goの
  RAW文字列リテラルによる
  複数行に渡る
  文字列
  `
	fmt.Printf("%v", s)

	// array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)
}
