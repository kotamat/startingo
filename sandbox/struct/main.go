package strct

import "fmt"

type User struct {
	Time
	Name  string
	Email string
}

type Admin struct {
	Time
	Account string
	Pass    string
}

func main() {
	u := User{
		Time{UpdatedAt: "now", CreatedAt: "now"},
		Name:  "name",
		Email: "kota1861@gmail.com",
	}

	a := Admin{
		Time:    {UpdatedAt: "yesterday", CreatedAt: "yesterday"},
		Account: "admin",
		Pass:    "pass",
	}
	fmt.Println(u, a)
}
