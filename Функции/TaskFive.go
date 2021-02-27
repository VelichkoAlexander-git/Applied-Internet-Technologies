package main

import "fmt"

func main() {
	i := fib(5)
	fmt.Println(i)
}

func fib(n uint) uint {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
