// Create a program that shows the type of some variable (use fmt.Printf)

package main

import "fmt"

func main() {
	a := 1
	b := 12.2
	c := "hey!"

	fmt.Printf("The value of a is: %T.\n", a)
	fmt.Printf("The value of b is: %T.\n", b)
	fmt.Printf("The value of c is: %T.\n", c)

}
