package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

//WriteToConsole takes a handler and returns a handler
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

//Added necessary middleware to save and load Session

//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next) // creates a new handler

	// set values vie cookie
	csrfHandler.SetBaseCookie(http.Cookie{ //token needs to be available on a per page basis
		HttpOnly: true,
		Path:     "/",              // apply to entire sites
		Secure:   app.InProduction, // https or not
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {

	//LoadAndSave automatically loads and saves session data for the current request and
	//communicates the session token to and from the client in a cookie
	return session.LoadAndSave(next)
}
