package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("TEMP DIR:", os.TempDir())
	http.ListenAndServe(":8080", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			sourc, _, err := req.FormFile("my-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer sourc.Close()

			dest, err := os.Create(filepath.Join(os.TempDir(), "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dest.Close()

			io.Copy(dest, sourc)
		}

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `
      <form method="POST" enctype="multipart/form-data">
        <input type="file" name="my-file">
        <input type="submit">
      </form>
      `)
	}))
}
