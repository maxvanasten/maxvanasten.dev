package main

import (
	"html/template"
	"net/http"
)

func GetRoutes() map[string]http.HandlerFunc {
	Routes := make(map[string]http.HandlerFunc)

	Routes["/"] = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("pages/index.html"))

		tmpl.Execute(w, nil)
	}

	Routes["/introduction"] = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("pages/partials/introduction.html"))

		tmpl.Execute(w, nil)
	}

	return Routes
}
