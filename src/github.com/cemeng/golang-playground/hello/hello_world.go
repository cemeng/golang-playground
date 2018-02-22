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

func showNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	nameFromPost := r.FormValue("name")
	vars := mux.Vars(r)
	nameFromMux := vars["name"]

	fmt.Fprintf(w, "Say my name, get name: %s, post name: %s, mux name: %s", name, nameFromPost, nameFromMux)
}

func createNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	nameFromPost := r.FormValue("name")
	fmt.Fprintf(w, "Say my name, get name: %s, post name: %s", name, nameFromPost)
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/names/{name}", showNameHandler).Methods("GET")
	r.HandleFunc("/names/", createNameHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
