package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var tpl *template.Template

const gcsBucket = "team_project_bucket"

func init() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img/"))))
	tpl = template.Must(template.ParseGlob("*.html"))
}

func handler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	tpl.ExecuteTemplate(res, "index.html", nil)

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {

		mpf, hdr, err := req.FormFile("dahui")
		if err != nil {
			log.Errorf(ctx, "ERROR handler req.FormFile: ", err)
			http.Error(res, fmt.Sprint("Error retrieving file", mpf, hdr, err), http.StatusInternalServerError)
			return
		}
		defer mpf.Close()

		fname, err := uploadFile(req, mpf, hdr)
		if err != nil {
			log.Errorf(ctx, "ERROR handler uploadFile: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		fnames, err := putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		html += `<h1>Files</h1>`
		for k, _ := range fnames {
			html += `<h3>` + k + `</h3>`
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
	tpl.ExecuteTemplate(res, "index1.html", nil)
}
