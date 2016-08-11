package main

import "fmt"

type Base struct {
	Id    int
	Owner string
}

type A struct {
	Base
	Name string
	Area string
}

type B struct {
	Base
	Title  string
	Bodies []string
}

func main() {
	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}
	b := B{
		Base:   Base{81, "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	fmt.Println(a.Id)
	fmt.Println(a.Owner)
	fmt.Println(b.Id)
	fmt.Println(b.Owner)

}
