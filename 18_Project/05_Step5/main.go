package main

import (
	"crypto/sha256"
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"io"
	"net/http"
	"text/template"
)

type User struct {
	Age  string
	Name string
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func anything(res http.ResponseWriter, req *http.Request) {
	user := User{}

	temp, err := template.ParseFiles("templates/temps.html")
	if err != nil {
		log.Fatalln(err)
	}
	temp.Execute(res, nil)

	cookie, err := req.Cookie("session-fino")

	user.Age = req.FormValue("Age")
	user.Name = req.FormValue("Name")

	UserData, err := json.Marshal(user)
	b64 := base64.URLEncoding.EncodeToString(UserData)

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "-" + user.Name + "-" + user.Age + "-" + getCode(b64),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", anything)
	http.ListenAndServe(":8080", nil)
}
