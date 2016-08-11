package main

import "fmt"

type User struct {
	Id   int
	Name string
}

// NewUser Userのコンストラクタ。
// New*とし、ポインターを返り値にするのが一般的
func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u

}
func main() {
	fmt.Println(NewUser(1, "Taro"))
}
