// Create a template that uses conditional logic

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmp, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
