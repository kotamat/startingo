package main

import "fmt"

func main() {

	cc()

	verifyClose()
}

func cc() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func verifyClose() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	var (
		i  int
		ok bool
	)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)

}
