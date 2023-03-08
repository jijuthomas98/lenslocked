package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome Jiju Thomas, modd is  </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch mail me at <a href=\"mailto: jijuthomas@email.com\">jijuthomas@email.com</a>")
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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHadler(w, r)
	default:
		http.Error(w, "page not found error", http.StatusNotFound)
	}
}

func main() {
	var router Router

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", router)

}
