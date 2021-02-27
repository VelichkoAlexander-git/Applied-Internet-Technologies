package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x, y, a, b float64
}

type Circle struct {
	x, y, r float64
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.a + r.b)
}

func (r *Rectangle) area() float64 {
	return r.a * r.b
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	area() float64
	perimeter() float64
}

func Calculation(f func(Shape) float64, shapes ...Shape) []float64 {
	array := make([]float64, len(shapes))
	for i, s := range shapes {
		array[i] = f(s)
	}
	return array
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x: 0, y: 0, a: 5, b: 3}
	value := Calculation(func(s Shape) float64 { return s.area() }, &c, &r)
	fmt.Println(value)
	value = Calculation(func(s Shape) float64 { return s.perimeter() }, &c, &r)
	fmt.Println(value)
}
