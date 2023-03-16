package main

import (
	"fmt"
	"test/project/circle"
	"test/project/rect"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := rect.Rect{Width: 3, Height: 4}
	c := circle.Circle{Radius: 5}

	measure(r)
	measure(c)
}
