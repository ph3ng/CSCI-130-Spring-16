package entrychecker

import "html/template"

import "net/http"

var tpl *template.Template

func init() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/api/check", checker)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
	tpl = template.Must(template.ParseGlob("*.html"))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}

// func checker(res http.ResponseWriter, req *http.Request) {
// 	ctx := appengine.NewContext(req)
//
// 	bs, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		http.Error(res, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	log.Infof(ctx, "received %v", string(bs))
// 	io.WriteString(res, "we received"+string(bs))
// }
