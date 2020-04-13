package main

import (
	"fmt"
)

type nums []int

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return !isEven(n)
}

func main() {
	numbers := nums{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}

	for _, n := range numbers {
		if isEven(n) {
			fmt.Println("even")
		}

		if isOdd(n) {
			fmt.Println("odd")
		}
	}
}