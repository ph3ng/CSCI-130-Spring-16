// Write a function with one variadic parameter that finds the greatest
// number in a list of numbers.

package main

import "fmt"

func greatest(num ...int) int {
	var max int
	for _, i := range num {
		if max < i {
			max = i
		}
	}
	return max
}

func main() {
	largest := greatest(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(largest)
}
