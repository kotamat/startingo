package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

func NewUser(id int, name string, age uint) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	u.Age = age
	return u
}

func main() {
	u := NewUser(1, "Taro", 32)
	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
}
