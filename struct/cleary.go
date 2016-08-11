package main

import "fmt"

type Point struct{ X, Y int }
type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}
func main() {
	ps := Points{}
	ps = append(ps, &Point{1, 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{3, 4})
	fmt.Println(ps.ToString())
}
