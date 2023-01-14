package main

import (
	"fmt"
	"github.com/KoosieDoosie/newwebsite/pkg/config"
	"github.com/justinas/nosurf"
	"net/http"
)

var app config.AppConfig

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r) // middleware program
	})
}

// adds crsf protection to all post requests
func NoSurf(next http.Handler) http.Handler {

	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler
}

// load and save the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
