package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/marif226/basic-webapp/pkg/config"
	"github.com/marif226/basic-webapp/pkg/handlers"
	"github.com/marif226/basic-webapp/pkg/render"
)

const portNumber = ":8080"

// Main application function
func main() {
	var app config.AppConfig

	// create template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	}

	// store template cache in application
	app.TemplateCache = templateCache
	app.UseCache = false

	// create new repository that holds app config
	repo := handlers.NewRepo(&app)
	// set this repository for handlers package
	handlers.NewHandlers(repo)

	// set app config for render package
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}