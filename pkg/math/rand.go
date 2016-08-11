package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.Seed(time.Now().UnixNano())
	fmt.Println(r)
	var a string

	fmt.Print(a)
}
