package main

import "fmt"

func main() {
	words := map[string]string{
		"hush":   "be quiet",
		"0":      "a number; a loser; emptiness; enlightenment",
		"golang": "the best programming language ever",
	}
	words["this is"] = "practice!"
	delete(words, "this is")
	words["this is"] = "practice!"
	words["0"] = "off"
	for k, v := range words {
		fmt.Println(k, v)
		//fmt.Println(v)
	}

}
