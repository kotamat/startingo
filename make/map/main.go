package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	// overwrite
	m2 := make(map[string]string)

	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	m2["Yamada"] = "Jiro"

	fmt.Println(m2)
}
