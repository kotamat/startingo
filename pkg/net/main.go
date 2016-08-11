package main

import (
	"fmt"
	"log"

	"net/http"
	"net/url"
)

func main() {

	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("https://example.com", vs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
