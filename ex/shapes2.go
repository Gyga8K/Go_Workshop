package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Circumference
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Geometry interface {
	Area() float64
	Perimeter() float64
}

func Measure(g Geometry) {
	area := g.Area()
	fmt.Println("Area:", area)
	perimeter := g.Perimeter()
	fmt.Println("Perimeter:", perimeter)
}
func main() {
	shape := os.Args[1]
	if shape == "circle" {
		r := os.Args[2]
		radius, _ := strconv.Atoi(r)
		circle := Circle{float64(radius)}
		Measure(circle)
	} else {
		w := os.Args[2]
		h := os.Args[3]
		width, _ := strconv.Atoi(w)
		height, _ := strconv.Atoi(h)
		rectangle := Rectangle{float64(width), float64(height)}
		Measure(rectangle)
	}
}
