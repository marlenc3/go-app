package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my contact!</h1><p>To get in touch, email at <a href=\"mailto:marelncarrillo3@gmail.com\">marlencarrillo3@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
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
</ul>`)
}

//func pathHanlder(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/contact":
//		contactHandler(w, r)
//	case "/":
//		homeHandler(w, r)
//	default:
//		http.Error(w, "Page not found", http.StatusNotFound)
//	}
//}

// type Router struct{}
//
//	func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		switch r.URL.Path {
//		case "/contact":
//			contactHandler(w, r)
//		case "/faq":
//			faqHandler(w, r)
//		case "/":
//			homeHandler(w, r)
//		default:
//			http.Error(w, "Page not found", http.StatusNotFound)
//		}
//	}
func getContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to "+id+" contact!</h1><p>To get in touch, email at <a href=\"mailto:marelncarrillo3@gmail.com\">marlencarrillo3@gmail.com</a></p>")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/", homeHandler)

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/{id}", getContact)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("starting server on :3000....")
	http.ListenAndServe(":3000", r)
}
