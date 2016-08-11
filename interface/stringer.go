package main

import "fmt"

type T struct {
	Id   int
	Name string
}

// Stringerインターフェース
// fmt.Println()した時にフォーマットする
func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)

}

func main() {
	t := &T{1, "taro"}
	fmt.Println(t)

}
