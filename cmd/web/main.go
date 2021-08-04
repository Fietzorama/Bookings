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

// split up files and new command "go run cmd/web/*.go"
const portNumber = ":8080" // can't be changed
var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	//Initialization of the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // how long lasts the session
	//Using cookie to store session
	session.Cookie.Persist = true                  //set wether cookie persist after window is closed by the user
	session.Cookie.SameSite = http.SameSiteLaxMode //what site the cookie the applies to
	session.Cookie.Secure = app.InProduction       //insist that the cookie are encrypted (https)

	//store session in application wide configuration
	app.Session = session

	// go and get the template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false //before declare repository, app use cache = false, see why in renders first if loop

	repo := handlers.NewRepo(&app) //create repo variable
	handlers.NewHandlers(repo)     //pass back to handlers and create new handlers
	render.NewTemplates(&app)      //give app render component(or render package) of app access to app config variable

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{ //starting server
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
