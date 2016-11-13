package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("run at localhost:8912")
	http.HandleFunc("/sw-test/", func(w http.ResponseWriter, r *http.Request) {
		filename := strings.TrimPrefix(r.URL.Path, "/sw-test/")
		f, err := os.Open("static/" + filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		defer f.Close()
		if strings.HasSuffix(filename, ".js") {
			w.Header().Add("Content-Type", "application/javascript")
		} else if strings.HasSuffix(filename, ".jpg") {
			w.Header().Add("Content-Type", "image/jpeg")
		}
		io.Copy(w, f)
		fmt.Println(r.Method, r.URL)
	})
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8912", nil)
}
