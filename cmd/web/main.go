package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/marif226/basic-webapp/pkg/config"
	"github.com/marif226/basic-webapp/pkg/handlers"
	"github.com/marif226/basic-webapp/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// Main application function
func main() {
	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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

	serv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting application on port %s\n", portNumber)
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}