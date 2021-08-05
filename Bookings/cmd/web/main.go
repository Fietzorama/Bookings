package main

import (
	"fmt"
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/handlers"
	"github.com/Fietzorama/Bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// it imports ./pkg/config; access to all content in config.go
// AppConfig is struct that holds things to share with our application
var app config.AppConfig

//  create var to make session variable avaible in my middleware etc
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false //:8080 is not encrypted connection

	//set up session
	session = scs.New()               // create a new session
	session.Lifetime = 24 * time.Hour //how long the session lives (10min. in high security)
	session.Cookie.Persist = true     // Session persist after browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // encrypted cookie and https connection, default in production

	app.Session = session

	// we assign var app and then we get the TemplateCache
	// its stored in the package render in the func CreateTemplate Cache
	tc, err := render.CreateTemplateCache()
	// if cannot be loaded
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// Now we got the TemplateCache lets start the app
	// we assign the TemplateCache to var app
	app.TemplateCache = tc

	// assigned the value of false to use cache
	// change in hmtl template code should be visible after editing
	// because we don't read the code from the disk anymore.
	// pulling it from that template cache that created
	app.UseCache = false // if change in hmtl template

	// create the Repository variable that calls New.Repo
	// NewRepo(handlers.go) says take this app config that you just passed,
	// appointed to the app config and populate the struct repository,
	// return a new instance of this type that holds the application.
	repo := handlers.NewRepo(&app)

	// after creating, pass it back and create new handlers
	// I call NewHandlers and NewHandlers actually sets that variable repo
	// but I'm not using that anywhere.
	// Change: HandlerFunc Home/About + "m *Repository" all handlers have access to repo
	handlers.NewHandlers(repo)

	// calling the new template cache
	// func NewTemplate in ./render returns a pointer
	// give our application the render package of application
	// access to app config variable
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// instead of _=http.ListenAndServe(pportNumber, nil)
	//The HTTP Server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// HTTP service running in this program as well. The valve context is set
	// as a base context on the server listener at the point where we instantiate
	// the server - look lower.
	// Run the server
	err = srv.ListenAndServe()
	log.Fatal(err)

}
