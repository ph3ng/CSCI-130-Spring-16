package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", anything)
	http.ListenAndServe(":8080", nil)
}

func anything(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("templates/temps.html")
	if err != nil {
		log.Fatalln(err)
	}
	temp.Execute(res, nil)

	cookie, err := req.Cookie("session-fino")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

}
