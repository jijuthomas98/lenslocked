package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tml, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was error parsing template.", http.StatusInternalServerError)
		return
	}
	data := ViewData{Name: "Jiju", Widgets: []Widget{
		{"Blue Widget", 12},
		{"Red Widget", 12},
		{"Green Widget", 12},
	}}
	err = tml.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was error executing template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func faqHadler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHadler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)

}
