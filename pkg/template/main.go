package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "name1"}
	f2 := Friend{Fname: "name2"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`
    hello {{.UserName}}!
    {{range .Emails}}
      an email {{.}}
    {{end}}
    {{with .Friends}}
    {{range .}}
      my friend name is {{.Fname}}
    {{end}}
    {{end}}
    `)
	p := Person{
		UserName: "Kota",
		Emails:   []string{"kota-matsumoto@altplus.co.jp", "kota1861@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}
