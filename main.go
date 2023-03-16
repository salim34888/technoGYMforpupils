package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, "data goes here")
}

func docs(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/docs.html"))
	tmpl.Execute(w, "data goes here")
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	tmpl.Execute(w, "data goes here")
}

func including(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/including.html"))
	tmpl.Execute(w, "data goes here")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/postupaushim", including)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/docs", docs)
	mux.HandleFunc("/", homePage)

	fileServer := http.FileServer(http.Dir("./templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
