package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
		"black": "#000000",
	}

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, value := range colors {
		fmt.Println(color, value)
	}
}