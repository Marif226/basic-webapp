package handlers

import (
	"net/http"
	"github.com/marif226/basic-webapp/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, t *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}