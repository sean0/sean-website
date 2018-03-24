package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
)

var views = template.Must(template.ParseGlob("*.html"))

func main() {
	addr := ":" + os.Getenv("PORT")
	s := pat.New()
	s.Get("/", index)
	log.Fatal(http.ListenAndServe(addr, s))
}

func index(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Cookie("name")

	w.Header().Set("Content-Type", "text/html")

	views.ExecuteTemplate(w, "index.html", struct {
		Name *http.Cookie
	}{
		Name: name,
	})
}
