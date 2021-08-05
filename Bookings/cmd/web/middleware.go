package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// General:
// middleware allows to process requests as they come in
// and make decisions about what to do with them
// Example, some pages on site that I only want logged in users to see.

//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	//token needs to be available on a per_page basis
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
// Web servers by their very nature are not state
// middleware that tells web server remember state using sessions
// LoadAndSave provides middleware, which automatically loads
// and saves session data for the current request
//and communicates the session token to and from the client in a cookie
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
