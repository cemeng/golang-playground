package main

import (
	"fmt"
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

func main() {
	http.HandleFunc("/", rootHandler)

	// how do I indicate if a route should be accessible by POST or GET only?
	// by adding Methods call at the end
	http.HandleFunc("/say", sayHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// r.URL.Path

	log.Fatal(http.ListenAndServe(":8000", nil))
}
