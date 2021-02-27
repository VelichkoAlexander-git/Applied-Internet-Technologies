package main

import "fmt"

func main() {
	i, flag := half(2)
	fmt.Println(i)
	fmt.Println(flag)
}

func half(number int) (int, bool) {
	if number%2 == 0 {
		return 0, true
	} else {
		return 1, false
	}
}
