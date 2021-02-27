package main

import (
	"fmt"
	"golang-book/chapter11/math"
)

func main() {
	xs := []float64{6, 10, 2, 7}
	avg := math.Average(xs)
	fmt.Println(avg)
	avg, _ = math.Min(xs)
	fmt.Println(avg)
	avg, _ = math.Max(xs)
	fmt.Println(avg)
}
