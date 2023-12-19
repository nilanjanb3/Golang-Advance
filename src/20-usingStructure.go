package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (r Rectangle) area() float64 {
	return r.length * r.width
}

func main() {
	c1 := Circle{radius: 7.0}
	r1 := Rectangle{length: 10.0, width: 20.0}

	var shapes []Shape = []Shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("Area: %v Perimeter: %v\n", shape.area(), shape.perimeter())

	}

	// var arr []int = []int{1, 2, 3}
	//or
	// arr := []int{1, 2, 3}
	//or
	// var arr []int
	// arr = []int{1, 2, 3, 4, 5}
}
