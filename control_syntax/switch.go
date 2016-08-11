package main

import "fmt"

func main() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")

	}

	// type
	var x interface{} = 3.4329
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type!")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("%#v\n", v)
	}
}
