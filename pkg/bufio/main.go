package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func normal() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー", err)
	}

}

func str() {
	s := `XXXX
  YYY
  ZZZ`

	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())

}
func main() {
	// normal()

	str()
}
