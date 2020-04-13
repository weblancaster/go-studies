package main

import (
	"fmt"
)

type address struct {
	city string
	zipCode int
}

type person struct {
	name string
	address address
}

func main() {
	jon := person{
		name: "jon doe",
		address: address{
			city: "costa mesa",
			zipCode: 92627,
		},
	}

	fmt.Println("before", jon)

	jon.updateName("jon don")

	fmt.Println("after", jon)
}

func (p *person) updateName(newName string) {
	(*p).name = newName
}