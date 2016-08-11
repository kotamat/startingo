package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t := time.Now()
	t = t.Add(d)
	fmt.Println(t)

	tt, err := time.Parse("2006 1 2", "2004 3 5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tt)
}
