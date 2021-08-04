package main

// using external package from github
// $ go get github.com/bmizerany/pat (check pwd!)
// check afterwards in go.mod

import (
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	//create a multiplexer (mux) - what is basically a http handler
	mux := chi.NewRouter()

	//install middleware, process on some request and perform some action on it
	//app throws panic and will die, instead get usefull information what went wrong and i wanna recover from that panic
	//Implement BEFORE the routing
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./static/")) // Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static/*", fileServer))
	return mux

}
