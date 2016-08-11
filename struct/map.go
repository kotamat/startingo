package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {

	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jiro"}: "Osaka",
	}

	fmt.Println(m1)

	m2 := map[int]User{
		1: {Id: 1, Name: "Taro"},
		2: {Id: 2, Name: "Jiro"},
	}

	fmt.Println(m2)
}
