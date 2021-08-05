package main

import (
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

/* GENERAL:
// When the request reaches the server, a multiplexer will inspect the URL being requested
// and redirect the request to the correct handler.
// Once the request reaches a handler, the handler will retrieve information from the request
// and process it accordingly.
// When the processing is complete, the handler passes the data to the template engine,
// which will use templates to generate HTML to be returned to the client.
*/

//routes receives a pointer to config.AppConfig and return a HTTP Handler
func routes(app *config.AppConfig) http.Handler {

	// ChiRouter(3rdParty) is a multiplexer (a handler) to implement the ServeHTTP
	// multiplexer redirects HTTP requests to the correct handler for processing
	// including static files.
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(middleware.Recoverer) // Gracefully absorbs panics and prints the stack trace
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// request reaches the server, a multiplexer will inspect the URL being requested
	// and redirect the request to the correct handler
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// fileServer create a handler which responds to all HTTP requests with the contents of a given file system.
	// For our file system we're using the "static" directory relative to our application,
	// but you could use any other directory on our machine
	fileServer := http.FileServer(http.Dir("./static/"))

	// use the http.Handle() function to register the file server as the handler for all requests,
	// and launch the server listening on port 8080.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
