package main

import (
	"fmt"
	"time"
)

func receiver(ch1 <-chan int, ch2 <-chan int) {
	for {
		select {
		case e1, ok := <-ch1:
			if ok == false {
				break
			}
			fmt.Println(e1)
		case e2, ok := <-ch2:
			if ok == false {
				break
			}
			fmt.Println(e2)
		default:
			fmt.Println("none")

		}

	}

}
func main() {

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go receiver(ch1, ch2)

	i := 0
	for i < 20 {
		ch1 <- i
		ch2 <- i
		i++
	}

	close(ch1)
	close(ch2)

	time.Sleep(3 * time.Second)
}
