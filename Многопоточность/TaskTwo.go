package main

import (
	"fmt"
	"time"
)

func Sleep(sec int) {
	ch := time.After(time.Second * time.Duration(sec))
	for {
		select {
		case <-ch:
			fmt.Println("slept!")
			return
		}
	}
}

func main() {
	go Sleep(10)

	var input string
	fmt.Scanln(&input)
}
