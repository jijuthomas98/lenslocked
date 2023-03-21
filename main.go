package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jijuthomas98/lenslocked/controllers"
	"github.com/jijuthomas98/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)

}
