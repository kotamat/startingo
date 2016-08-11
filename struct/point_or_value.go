package main

import "fmt"

type Point struct {
	X, Y int
}

//Point型のレシーバー => 値は更新されない
// func (p Point) Set(x, y int) {
//
// 	p.X = x
// 	p.Y = y
// }

// *Point型のレシーバー => 値は更新される
func (p *Point) Set(x, y int) {

	p.X = x
	p.Y = y
}

func main() {
	p1 := Point{}
	p1.Set(1, 2)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p2 := &Point{}
	p2.Set(1, 2)
	fmt.Println(p2.X)
	fmt.Println(p2.Y)

}
