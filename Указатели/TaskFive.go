package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(x)
	fmt.Println(y)
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
