package main

import "fmt"

func main() {

	var (
		ch0 chan int   // 送受信 他の専用の方に代入可能
		ch1 <-chan int // 受信専用
		ch2 chan<- int // 送信専用
	)
	ch0 = ch1 // NG
	ch0 = ch2 // NG
	ch1 = ch0
	ch1 = ch2 // NG
	ch2 = ch0
	ch2 = ch1 // NG

	fmt.Println(ch0)
	fmt.Println(ch1)
	fmt.Println(ch2)
}
