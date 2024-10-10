package main

import (
	"fmt"
	"net/http"
	"html/template"
)
func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("about.html")
	t.Execute(w,p)
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}