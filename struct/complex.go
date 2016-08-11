package main

import "fmt"

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed // フィールド名を省略した構造体の埋め込み
}

func main() {

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a.Amount) // => 10

	a.Amount = 15
	fmt.Println(a.Amount)      // => 15
	fmt.Println(a.Feed.Amount) // => 15
}
