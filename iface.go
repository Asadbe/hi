package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perimetr() float32
}
type rectangle struct {
	width, height float32
}
type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}
func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) perimetr() float32 {
	return 2 * (r.width + r.height)
}
func (c circle) perimetr() float32 {
	return math.Pi * c.radius * 2
}
func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimetr())
}
func main() {
	rec := rectangle{width: 5, height: 2}
	cir := circle{radius: 10}
	measure(rec)
	measure(cir)

}
