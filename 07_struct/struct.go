package main

import "fmt"

type person struct {
	name string
	pets []string
}

func main() {
	x := person{
		name: "Pheng",
		pets: []string{"Mochi"},
	}

	fmt.Println(x)
}
