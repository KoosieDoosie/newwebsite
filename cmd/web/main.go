package main

import (
	"fmt"
	"github.com/KoosieDoosie/newwebsite/pkg/config"
	"github.com/KoosieDoosie/newwebsite/pkg/handlers"
	"github.com/KoosieDoosie/newwebsite/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var session *scs.SessionManager

// main is the main function
func main() {
	var app config.AppConfig

	//Change this to ture when ready to go in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //How long the session will stay open
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(" Cannot create the template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app) //give access to the app var

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
