package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tml, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was error parsing template.", http.StatusInternalServerError)
		return
	}
	err = tml.Execute(w, "someName")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	  <li>
		<b>Is there a free version?</b>
		Yes! We offer a free trial for 30 days on any paid plans.
	  </li>
	  <li>
		<b>What are your support hours?</b>
		We have support staff answering emails 24/7, though response
		times may be a bit slower on weekends.
	  </li>
	  <li>
		<b>How do I contact support?</b>
		Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	  </li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHadler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)

}
