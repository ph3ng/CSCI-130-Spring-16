package main

import (
	"log"
	"net/http"
	"text/template"
)

func anything(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("templates/temps.html")
	logFatalError(err)
	temp.Execute(res, nil)
	logFatalError(err)
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", anything)
	http.ListenAndServe(":8080", nil)
}

// Logs error at Fatal level given the error is not nil.
func logFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
