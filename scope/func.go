package main

import "fmt"

func appName() string {
	const AppName = "Go Application"
	var Version = "1.0"
	return AppName + " " + Version
}

func main() {
	fmt.Println(appName())
	fmt.Println(AppName + " " + Version)
}
