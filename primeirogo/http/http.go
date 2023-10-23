package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!!"))
}

var templates *template.Template

func htmlTemplate(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}
func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)
	http.HandleFunc("/html", htmlTemplate)

	fmt.Println("Rodando na porta 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
