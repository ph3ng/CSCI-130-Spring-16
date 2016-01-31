//Create the ""surfer page"" and serve it using Go

package main
import (
      "html/template"
      "net/http"
      "log"
)

func surfer(res http.ResponseWriter, req *http.Request)  {
    tmp, err := template.ParseFiles("./index.html")
    if err != nil {
        log.Fatalln(err)
    }
    tmp.Execute(res, nil)
}

/*func img(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "pics/surf.jpg")
}

func back(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "css/main.css")
}
*/

func main() {
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.HandleFunc("/", surfer)
    http.ListenAndServe(":8080", nil)
}
