package main

import "fmt"

func main() {
	i := max(2, 5, 2, 6)
	fmt.Println(i)
}

func max(args ...int) int {
	max := args[0]
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}
