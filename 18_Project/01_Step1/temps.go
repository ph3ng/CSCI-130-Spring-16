package main

import (
	"log"
	"net/http"
	"text/template"
)

func anything(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("templates/temps.html")
	if err != nil {
		log.Fatalln(err)
	}
	temp.Execute(res, nil)
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", anything)
	http.ListenAndServe(":8080", nil)
}
