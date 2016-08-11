package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// testRecover
	testRecover(839201)

	panic("panic!")

	fmt.Println("Hello World")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")

			}
		}
	}()

	panic(src)
	return
}
