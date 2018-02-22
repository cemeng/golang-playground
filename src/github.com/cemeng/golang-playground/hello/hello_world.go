package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome this is the root")
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	nameFromPost := r.FormValue("name")
	fmt.Fprintf(w, "Say my name, get name: %s, post name: %s", name, nameFromPost)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/say", sayHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/books/{title}/page/{page}", booksHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
